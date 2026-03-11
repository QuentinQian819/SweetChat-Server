package user

import (
	"chatBox/api/v1"
	"chatBox/internal/logic/user"
	"context"

	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

// Register 用户注册
func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	return user.User().Register(ctx, req)
}

// Login 用户登录
func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	return user.User().Login(ctx, req)
}

// GenerateInvite 生成情侣邀请码
func (c *ControllerV1) GenerateInvite(ctx context.Context, req *v1.GenerateInviteReq) (res *v1.GenerateInviteRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return user.User().GenerateInvite(ctx, userId)
}

// BindCouple 绑定情侣关系
func (c *ControllerV1) BindCouple(ctx context.Context, req *v1.BindCoupleReq) (res *v1.BindCoupleRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return user.User().BindCouple(ctx, userId, req.InviteCode)
}

// GetProfile 获取个人信息
func (c *ControllerV1) GetProfile(ctx context.Context, req *v1.GetProfileReq) (res *v1.GetProfileRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return user.User().GetProfile(ctx, userId)
}

// UpdateProfile 更新个人信息
func (c *ControllerV1) UpdateProfile(ctx context.Context, req *v1.UpdateProfileReq) (res *v1.UpdateProfileRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return user.User().UpdateProfile(ctx, userId, req)
}

// GetCoupleInfo 获取情侣信息
func (c *ControllerV1) GetCoupleInfo(ctx context.Context, req *v1.GetCoupleInfoReq) (res *v1.GetCoupleInfoRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return user.User().GetCoupleInfo(ctx, userId)
}
