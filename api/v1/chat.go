package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// GetHistoryReq 获取聊天历史请求
type GetHistoryReq struct {
	g.Meta   `path:"/chat/history" method:"get" tags:"聊天" summary:"获取聊天历史"`
	Page     int64  `json:"page" v:"min:1#页码最小为1"`
	PageSize int64  `json:"pageSize" v:"between:1,100#每页数量在1-100之间"`
	LastId   uint64 `json:"lastId"` // 分页用，获取此ID之前的消息
}

type GetHistoryRes struct {
	List        []*MessageItem `json:"list"`
	HasMore     bool           `json:"hasMore"`
	UnreadCount int64          `json:"unreadCount"`
}

type MessageItem struct {
	Id         uint64    `json:"id"`
	CoupleId   uint64    `json:"coupleId"`
	SenderId   uint64    `json:"senderId"`
	ReceiverId uint64    `json:"receiverId"`
	MsgType    int8      `json:"msgType"`
	Content    string    `json:"content"`
	IsRead     bool      `json:"isRead"`
	CreatedAt  time.Time `json:"createdAt"`
}

// MarkReadReq 标记消息已读请求
type MarkReadReq struct {
	g.Meta   `path:"/chat/read" method:"put" tags:"聊天" summary:"标记消息已读"`
	MessageId uint64 `json:"messageId" v:"required#请提供消息ID"`
}

type MarkReadRes struct {
	Success bool `json:"success"`
}

// UploadReq 上传文件请求 (multipart/form-data)
type UploadReq struct {
	g.Meta `path:"/chat/upload" method:"post" tags:"聊天" summary:"上传聊天文件"`
}

type UploadRes struct {
	Url      string `json:"url"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

// WebSocket message types
type WSMessage struct {
	Type string      `json:"type"` // message, ping, pong
	Data interface{} `json:"data"`
}

type WSChatMessage struct {
	Id         uint64 `json:"id"`
	CoupleId   uint64 `json:"coupleId"`
	SenderId   uint64 `json:"senderId"`
	ReceiverId uint64 `json:"receiverId"`
	MsgType    int8   `json:"msgType"`
	Content    string `json:"content"`
	IsRead     bool   `json:"isRead"`
	CreatedAt  int64  `json:"createdAt"`
}

// ClearReq 清空消息请求
type ClearReq struct {
	g.Meta `path:"/chat/clear" method:"post" tags:"聊天" summary:"清空所有消息"`
}

type ClearRes struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
