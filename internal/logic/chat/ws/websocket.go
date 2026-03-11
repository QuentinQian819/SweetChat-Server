package ws

import (
	"chatBox/internal/consts"
	"chatBox/internal/logic/chat"
	"chatBox/internal/logic/jwt"
	"context"
	"encoding/json"
	"fmt"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

// WebSocket message types from gorilla/websocket
const (
	WSMsgText  = 1  // TextMessage
	WSMsgBinary = 2
	WSMsgClose = 8
	WSMsgPing   = 9
	WSMsgPong   = 10
)

// WSManager WebSocket连接管理器
type WSManager struct {
	connections map[uint64]*WebSocketConn // userId -> websocket connection
	mu          sync.RWMutex
}

// WebSocketConn 封装的WebSocket连接
type WebSocketConn struct {
	UserId  uint64
	Phone   string
	Conn    *ghttp.WebSocket
}

var Manager = &WSManager{
	connections: make(map[uint64]*WebSocketConn),
}

// WSHandler WebSocket处理函数
func WSHandler(r *ghttp.Request) {
	ws, err := r.WebSocket()
	if err != nil {
		r.Response.WriteJson(g.Map{
			"code":    500,
			"message": "WebSocket upgrade failed",
		})
		return
	}

	// Initialize connection with empty user info (will be set after auth)
	currentConn := &WebSocketConn{
		UserId: 0,
		Phone:  "",
		Conn:   ws,
	}

	defer func() {
		// 连接关闭时清理
		if currentConn != nil && currentConn.UserId > 0 {
			Manager.RemoveConnection(currentConn.UserId)
			// 更新Redis离线状态
			g.Redis().Del(r.Context(), fmt.Sprintf(consts.OnlineUserPattern, currentConn.UserId))
		}
		ws.Close()
	}()

	for {
		msgType, message, err := ws.ReadMessage()
		if err != nil {
			return
		}

		// 处理消息
		if msgType == WSMsgText {
			Manager.HandleMessage(currentConn, message)
		}
	}
}

// AddConnection 添加连接
func (m *WSManager) AddConnection(userId uint64, phone string, ws *ghttp.WebSocket) *WebSocketConn {
	m.mu.Lock()
	defer m.mu.Unlock()

	conn := &WebSocketConn{
		UserId: userId,
		Phone:  phone,
		Conn:   ws,
	}
	m.connections[userId] = conn

	// 更新Redis在线状态
	_, err := g.Redis().Set(context.Background(), fmt.Sprintf(consts.OnlineUserPattern, userId), 1)
	if err != nil {
		g.Log().Errorf(context.Background(), "failed to set online status: %v", err)
	}

	return conn
}

// RemoveConnection 移除连接
func (m *WSManager) RemoveConnection(userId uint64) {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.connections, userId)
}

// GetConnection 获取连接
func (m *WSManager) GetConnection(userId uint64) (*WebSocketConn, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	conn, ok := m.connections[userId]
	return conn, ok
}

// SendMessageToUser 发送消息给指定用户
func (m *WSManager) SendMessageToUser(userId uint64, message interface{}) error {
	conn, ok := m.GetConnection(userId)
	if !ok {
		return fmt.Errorf("user %d not connected", userId)
	}

	data, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return conn.Conn.WriteMessage(WSMsgText, data)
}

// BroadcastToCouple 广播消息给情侣双方
func (m *WSManager) BroadcastToCouple(user1Id, user2Id uint64, message interface{}) {
	m.SendMessageToUser(user1Id, message)
	m.SendMessageToUser(user2Id, message)
}

// HandleMessage 处理接收到的消息
func (m *WSManager) HandleMessage(conn *WebSocketConn, data []byte) {
	var msg map[string]interface{}
	if err := json.Unmarshal(data, &msg); err != nil {
		return
	}

	msgType, _ := msg["type"].(string)

	switch msgType {
	case "auth":
		// 认证消息
		if conn == nil || conn.Conn == nil {
			return
		}

		token, _ := msg["token"].(string)
		claims, err := jwt.ParseToken(token)
		if err != nil {
			data, _ := json.Marshal(g.Map{"type": "error", "message": "认证失败"})
			conn.Conn.WriteMessage(WSMsgText, data)
			return
		}

		// 获取WebSocket连接
		ws := conn.Conn

		// 更新连接信息
		newConn := m.AddConnection(claims.UserId, claims.Phone, ws)

		// 更新引用
		// Hack: 更新外部引用
		conn.UserId = newConn.UserId
		conn.Phone = newConn.Phone

		data, _ := json.Marshal(g.Map{"type": "auth", "success": true})
		conn.Conn.WriteMessage(WSMsgText, data)

	case "message":
		// 聊天消息
		if conn == nil || conn.UserId == 0 {
			data, _ := json.Marshal(g.Map{"type": "error", "message": "请先认证"})
			if conn != nil && conn.Conn != nil {
				conn.Conn.WriteMessage(WSMsgText, data)
			}
			return
		}

		m.HandleChatMessage(conn, msg)

	case "ping":
		// 心跳
		if conn != nil && conn.Conn != nil {
			data, _ := json.Marshal(g.Map{"type": "pong"})
			conn.Conn.WriteMessage(WSMsgText, data)
		}
	}
}

// HandleChatMessage 处理聊天消息
func (m *WSManager) HandleChatMessage(conn *WebSocketConn, msg map[string]interface{}) {
	userId := conn.UserId

	// 获取情侣关系
	coupleInfo, err := chat.Chat().GetCoupleInfo(context.Background(), userId)
	if err != nil || coupleInfo == nil {
		data, _ := json.Marshal(g.Map{"type": "error", "message": "未绑定情侣关系"})
		if conn != nil && conn.Conn != nil {
			conn.Conn.WriteMessage(WSMsgText, data)
		}
		return
	}

	// 解析消息内容
	msgType := int8(gconv.Int(msg["msgType"]))
	content := gconv.String(msg["content"])

	if msgType == 0 {
		msgType = consts.MessageTypeText
	}

	// 保存消息到数据库
	db := g.DB()
	now := gtime.Now()
	messageId, err := db.Model("messages").Ctx(context.Background()).Data(g.Map{
		"couple_id":   coupleInfo.CoupleId,
		"sender_id":   userId,
		"receiver_id": coupleInfo.PartnerId,
		"msg_type":    msgType,
		"content":     content,
		"is_read":     0,
		"created_at":  now,
	}).InsertAndGetId()

	if err != nil {
		data, _ := json.Marshal(g.Map{"type": "error", "message": "消息发送失败"})
		if conn != nil && conn.Conn != nil {
			conn.Conn.WriteMessage(WSMsgText, data)
		}
		return
	}

	// 构造返回消息
	wsMsg := map[string]interface{}{
		"type": "message",
		"data": map[string]interface{}{
			"id":         messageId,
			"coupleId":   coupleInfo.CoupleId,
			"senderId":   userId,
			"receiverId": coupleInfo.PartnerId,
			"msgType":    msgType,
			"content":    content,
		},
	}

	// 发送给双方
	m.BroadcastToCouple(userId, coupleInfo.PartnerId, wsMsg)
}
