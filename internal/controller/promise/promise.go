package promise

import (
	"chatBox/api/v1"
	"chatBox/internal/logic/promise"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

// Create 创建承诺
func (c *ControllerV1) Create(ctx context.Context, req *v1.CreatePromiseReq) (res *v1.CreatePromiseRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return promise.Promise().Create(ctx, userId, req)
}

// List 获取承诺列表
func (c *ControllerV1) List(ctx context.Context, req *v1.PromiseListReq) (res *v1.PromiseListRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return promise.Promise().List(ctx, userId, req)
}

// Get 获取承诺详情
func (c *ControllerV1) Get(ctx context.Context, req *v1.GetPromiseReq) (res *v1.GetPromiseRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return promise.Promise().Get(ctx, userId, req.Id)
}

// Update 更新承诺
func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdatePromiseReq) (res *v1.UpdatePromiseRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return promise.Promise().Update(ctx, userId, req)
}

// Delete 删除承诺
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeletePromiseReq) (res *v1.DeletePromiseRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return promise.Promise().Delete(ctx, userId, req.Id)
}

// ToggleComplete 切换完成状态
func (c *ControllerV1) ToggleComplete(ctx context.Context, req *v1.ToggleCompletePromiseReq) (res *v1.ToggleCompletePromiseRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return promise.Promise().ToggleComplete(ctx, userId, req.Id)
}
