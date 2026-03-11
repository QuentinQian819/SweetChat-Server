package chat

import (
	"chatBox/api/v1"
	"chatBox/internal/logic/chat"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

// GetHistory 获取聊天历史
func (c *ControllerV1) GetHistory(ctx context.Context, req *v1.GetHistoryReq) (res *v1.GetHistoryRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return chat.Chat().GetHistory(ctx, userId, req)
}

// MarkRead 标记消息已读
func (c *ControllerV1) MarkRead(ctx context.Context, req *v1.MarkReadReq) (res *v1.MarkReadRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return chat.Chat().MarkRead(ctx, userId, req)
}

// Upload 上传文件
func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	file := r.GetUploadFile("file")
	if file == nil {
		return nil, gerror.New("请选择文件")
	}
	return chat.Chat().Upload(ctx, userId, file)
}

// Clear 清空所有消息
func (c *ControllerV1) Clear(ctx context.Context, req *v1.ClearReq) (res *v1.ClearRes, err error) {
	return chat.Chat().Clear(ctx)
}
