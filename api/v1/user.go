package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// RegisterReq 用户注册请求
type RegisterReq struct {
	g.Meta   `path:"/user/register" method:"post" tags:"用户" summary:"用户注册"`
	Phone    string `json:"phone" v:"required#请输入手机号"`
	Password string `json:"password" v:"required|length:6,32#请输入密码|密码长度为6-32位"`
	Nickname string `json:"nickname" v:"required#请输入昵称"`
}

type RegisterRes struct {
	UserId   uint64 `json:"userId"`
	Token    string `json:"token"`
	Nickname string `json:"nickname"`
}

// LoginReq 用户登录请求
type LoginReq struct {
	g.Meta   `path:"/user/login" method:"post" tags:"用户" summary:"用户登录"`
	Phone    string `json:"phone" v:"required#请输入手机号"`
	Password string `json:"password" v:"required#请输入密码"`
}

type LoginRes struct {
	UserId   uint64 `json:"userId"`
	Token    string `json:"token"`
	Nickname string `json:"nickname"`
	Avatar   string `json:"avatar"`
}

// GenerateInviteReq 生成邀请码请求
type GenerateInviteReq struct {
	g.Meta `path:"/user/generate-invite" method:"post" tags:"用户" summary:"生成情侣邀请码"`
}

type GenerateInviteRes struct {
	InviteCode string `json:"inviteCode"`
}

// BindCoupleReq 绑定情侣关系请求
type BindCoupleReq struct {
	g.Meta     `path:"/user/bind-couple" method:"post" tags:"用户" summary:"绑定情侣关系"`
	InviteCode string `json:"inviteCode" v:"required#请输入邀请码"`
}

type BindCoupleRes struct {
	CoupleId  uint64 `json:"coupleId"`
	PartnerId uint64 `json:"partnerId"`
}

// GetProfileReq 获取个人信息请求
type GetProfileReq struct {
	g.Meta `path:"/user/profile" method:"get" tags:"用户" summary:"获取个人信息"`
}

type GetProfileRes struct {
	UserId    uint64    `json:"userId"`
	Phone     string    `json:"phone"`
	Nickname  string    `json:"nickname"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"createdAt"`
}

// UpdateProfileReq 更新个人信息请求
type UpdateProfileReq struct {
	g.Meta   `path:"/user/profile" method:"put" tags:"用户" summary:"更新个人信息"`
	Nickname string `json:"nickname" v:"required#请输入昵称"`
	Avatar   string `json:"avatar"`
}

type UpdateProfileRes struct {
	UserId uint64 `json:"userId"`
}

// GetCoupleInfoReq 获取情侣信息请求
type GetCoupleInfoReq struct {
	g.Meta `path:"/user/couple-info" method:"get" tags:"用户" summary:"获取情侣信息"`
}

type GetCoupleInfoRes struct {
	CoupleId  uint64 `json:"coupleId"`
	UserId    uint64 `json:"userId"`    // 当前用户ID
	PartnerId uint64 `json:"partnerId"` // 伴侣ID
	Nickname  string `json:"nickname"`  // 伴侣昵称
	Avatar    string `json:"avatar"`    // 伴侣头像
}
