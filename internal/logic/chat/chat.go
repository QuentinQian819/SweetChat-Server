package chat

import (
	"chatBox/api/v1"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
)

type IChatLogic interface {
	GetHistory(ctx context.Context, userId uint64, in *v1.GetHistoryReq) (*v1.GetHistoryRes, error)
	MarkRead(ctx context.Context, userId uint64, in *v1.MarkReadReq) (*v1.MarkReadRes, error)
	Upload(ctx context.Context, userId uint64, file *ghttp.UploadFile) (*v1.UploadRes, error)
	GetCoupleInfo(ctx context.Context, userId uint64) (*CoupleInfo, error)
	Clear(ctx context.Context) (*v1.ClearRes, error)
}

type chatLogicImpl struct{}

type CoupleInfo struct {
	CoupleId  uint64
	UserId    uint64
	PartnerId uint64
}

func Chat() IChatLogic {
	return &chatLogicImpl{}
}

// GetHistory 获取聊天历史
func (s *chatLogicImpl) GetHistory(ctx context.Context, userId uint64, in *v1.GetHistoryReq) (*v1.GetHistoryRes, error) {
	db := g.DB()

	// 获取情侣关系
	coupleInfo, err := s.GetCoupleInfo(ctx, userId)
	if err != nil {
		return nil, err
	}
	if coupleInfo == nil {
		return nil, gerror.New("未绑定情侣关系")
	}

	// 默认分页参数
	page := in.Page
	if page <= 0 {
		page = 1
	}
	pageSize := in.PageSize
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	// 构建查询
	query := db.Model("messages").Ctx(ctx).
		Where("couple_id", coupleInfo.CoupleId)

	// 如果有lastId，查询更早的消息
	if in.LastId > 0 {
		query = query.WhereLT("id", in.LastId)
	}

	// 查询消息列表
	messages, err := query.OrderDesc("id").
		Limit(int(pageSize)).
		All()

	if err != nil {
		return nil, err
	}

	// 转换为返回格式
	list := make([]*v1.MessageItem, 0, len(messages))
	for _, msg := range messages {
		list = append(list, &v1.MessageItem{
			Id:         msg["id"].Uint64(),
			CoupleId:   msg["couple_id"].Uint64(),
			SenderId:   msg["sender_id"].Uint64(),
			ReceiverId: msg["receiver_id"].Uint64(),
			MsgType:    int8(msg["msg_type"].Int()),
			Content:    msg["content"].String(),
			IsRead:     msg["is_read"].Int() == 1,
			CreatedAt:  msg["created_at"].Time(),
		})
	}

	// 查询未读消息数
	unreadCount, err := db.Model("messages").Ctx(ctx).
		Where("couple_id", coupleInfo.CoupleId).
		Where("receiver_id", userId).
		Where("is_read", 0).
		Count()

	if err != nil {
		return nil, err
	}

	return &v1.GetHistoryRes{
		List:        list,
		HasMore:     len(list) == int(pageSize),
		UnreadCount: int64(unreadCount),
	}, nil
}

// MarkRead 标记消息已读
func (s *chatLogicImpl) MarkRead(ctx context.Context, userId uint64, in *v1.MarkReadReq) (*v1.MarkReadRes, error) {
	db := g.DB()

	// 更新消息已读状态
	_, err := db.Model("messages").Ctx(ctx).
		Where("id", in.MessageId).
		Where("receiver_id", userId).
		Update(g.Map{
		"is_read": 1,
	})

	if err != nil {
		return nil, err
	}

	return &v1.MarkReadRes{
		Success: true,
	}, nil
}

// Upload 上传文件
func (s *chatLogicImpl) Upload(ctx context.Context, userId uint64, file *ghttp.UploadFile) (*v1.UploadRes, error) {
	if file == nil {
		return nil, gerror.New("请选择文件")
	}

	// 检查文件大小
	config := g.Cfg()
	maxSize := config.MustGet(ctx, "upload.max_size").Int64()
	if file.Size > maxSize {
		return nil, gerror.Newf("文件大小不能超过%dMB", maxSize/1024/1024)
	}

	// 检查文件类型
	allowedTypes := config.MustGet(ctx, "upload.allowed_image_types").Strings()
	allowedTypes = append(allowedTypes, config.MustGet(ctx, "upload.allowed_audio_types").Strings()...)

	isAllowed := false
	contentType := file.Header.Get("Content-Type")
	for _, t := range allowedTypes {
		if strings.HasPrefix(contentType, t) {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return nil, gerror.New("不支持的文件类型")
	}

	// 生成文件名
	ext := filepath.Ext(file.Filename)
	newFilename := fmt.Sprintf("%s%s", guid.S(), ext)

	// 按日期分目录存储
	datePath := gtime.Now().Format("2006/01/02")
	relativePath := filepath.Join("uploads", datePath, newFilename)
	absolutePath := filepath.Join("resource", relativePath)

	// 确保目录存在
	dir := filepath.Dir(absolutePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0755)
	}

	// 保存文件
	savedName, err := file.Save(absolutePath, true)
	if err != nil {
		return nil, err
	}

	// 返回访问URL
	url := fmt.Sprintf("/resource/%s", filepath.Join("uploads", datePath, savedName))

	return &v1.UploadRes{
		Url:      url,
		Filename: file.Filename,
		Size:     file.Size,
	}, nil
}

// GetCoupleInfo 获取情侣信息
func (s *chatLogicImpl) GetCoupleInfo(ctx context.Context, userId uint64) (*CoupleInfo, error) {
	db := g.DB()

	couple, err := db.Model("couples").Ctx(ctx).
		Where("user1_id", userId).
		WhereOr("user2_id", userId).
		Where("status", 1).
		One()

	if err != nil {
		return nil, err
	}

	if couple.IsEmpty() {
		return nil, nil
	}

	coupleId := couple["id"].Uint64()
	user1Id := couple["user1_id"].Uint64()
	user2Id := couple["user2_id"].Uint64()

	var partnerId uint64
	if userId == user1Id {
		partnerId = user2Id
	} else {
		partnerId = user1Id
	}

	return &CoupleInfo{
		CoupleId:  coupleId,
		UserId:    userId,
		PartnerId: partnerId,
	}, nil
}

// Clear 清空所有消息
func (s *chatLogicImpl) Clear(ctx context.Context) (*v1.ClearRes, error) {
	_, err := g.DB().Model("messages").Ctx(ctx).Delete()
	if err != nil {
		return nil, err
	}

	return &v1.ClearRes{
		Success: true,
		Message: "清空成功",
	}, nil
}
