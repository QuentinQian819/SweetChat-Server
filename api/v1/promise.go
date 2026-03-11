package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// CreatePromiseReq 创建承诺请求
type CreatePromiseReq struct {
	g.Meta     `path:"/promise" method:"post" tags:"承诺" summary:"创建承诺"`
	Title      string   `json:"title" v:"required#请输入标题"`
	MessageIds []uint64 `json:"messageIds"` // 关联的消息ID列表
	ColorTag   int      `json:"colorTag"`   // 颜色标签
}

type CreatePromiseRes struct {
	PromiseId uint64 `json:"promiseId"`
}

// PromiseListReq 获取承诺列表请求
type PromiseListReq struct {
	g.Meta   `path:"/promise/list" method:"get" tags:"承诺" summary:"获取承诺列表"`
	Page     int64 `json:"page" v:"min:1#页码最小为1"`
	PageSize int64 `json:"pageSize" v:"between:1,100#每页数量在1-100之间"`
}

type PromiseListRes struct {
	List     []*PromiseItem `json:"list"`
	Total    int64          `json:"total"`
	Page     int64          `json:"page"`
	PageSize int64          `json:"pageSize"`
}

type PromiseItem struct {
	Id          uint64     `json:"id"`
	Title       string     `json:"title"`
	MessageIds  []uint64   `json:"messageIds"`
	ColorTag    int        `json:"colorTag"`
	IsCompleted bool       `json:"isCompleted"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`
	CreatedAt   time.Time  `json:"createdAt"`
	UpdatedAt   time.Time  `json:"updatedAt"`
}

// GetPromiseReq 获取承诺详情请求
type GetPromiseReq struct {
	g.Meta `path:"/promise/:id" method:"get" tags:"承诺" summary:"获取承诺详情"`
	Id     uint64 `json:"id" v:"required#请提供承诺ID"`
}

type GetPromiseRes struct {
	PromiseItem
}

// UpdatePromiseReq 更新承诺请求
type UpdatePromiseReq struct {
	g.Meta     `path:"/promise/:id" method:"put" tags:"承诺" summary:"更新承诺"`
	Id         uint64   `json:"id" v:"required#请提供承诺ID"`
	Title      string   `json:"title" v:"required#请输入标题"`
	MessageIds []uint64 `json:"messageIds"`
	ColorTag   int      `json:"colorTag"`
}

type UpdatePromiseRes struct {
	Success bool `json:"success"`
}

// DeletePromiseReq 删除承诺请求
type DeletePromiseReq struct {
	g.Meta `path:"/promise/:id" method:"delete" tags:"承诺" summary:"删除承诺"`
	Id     uint64 `json:"id" v:"required#请提供承诺ID"`
}

type DeletePromiseRes struct {
	Success bool `json:"success"`
}

// ToggleCompletePromiseReq 切换完成状态请求
type ToggleCompletePromiseReq struct {
	g.Meta `path:"/promise/:id/complete" method:"put" tags:"承诺" summary:"切换承诺完成状态"`
	Id     uint64 `json:"id" v:"required#请提供承诺ID"`
}

type ToggleCompletePromiseRes struct {
	Success     bool       `json:"success"`
	IsCompleted bool       `json:"isCompleted"`
	CompletedAt *time.Time `json:"completedAt,omitempty"`
}
