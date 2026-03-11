package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

// CreateDiaryReq 创建日记请求
type CreateDiaryReq struct {
	g.Meta   `path:"/diary" method:"post" tags:"日记" summary:"创建日记"`
	Title    string   `json:"title" v:"required#请输入标题"`
	Content  string   `json:"content" v:"required#请输入内容"`
	IsShared int8     `json:"isShared"` // 是否共享给对方
	Mood     string   `json:"mood"`     // 心情标签
	Weather  string   `json:"weather"`  // 天气
	MediaIds []uint64 `json:"mediaIds"` // 附件ID列表
}

type CreateDiaryRes struct {
	DiaryId uint64 `json:"diaryId"`
}

// DiaryListReq 获取日记列表请求
type DiaryListReq struct {
	g.Meta   `path:"/diary/list" method:"get" tags:"日记" summary:"获取日记列表"`
	Page     int64  `json:"page" v:"min:1#页码最小为1"`
	PageSize int64  `json:"pageSize" v:"between:1,100#每页数量在1-100之间"`
}

type DiaryListRes struct {
	List     []*DiaryItem `json:"list"`
	Total    int64        `json:"total"`
	Page     int64        `json:"page"`
	PageSize int64        `json:"pageSize"`
}

type DiaryItem struct {
	Id        uint64    `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	IsShared  bool      `json:"isShared"`
	Mood      string    `json:"mood"`
	Weather   string    `json:"weather"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Media     []*MediaItem `json:"media"`
}

// GetDiaryReq 获取日记详情请求
type GetDiaryReq struct {
	g.Meta `path:"/diary/:id" method:"get" tags:"日记" summary:"获取日记详情"`
	Id     uint64 `json:"id" v:"required#请提供日记ID"`
}

type GetDiaryRes struct {
	DiaryItem
}

// UpdateDiaryReq 更新日记请求
type UpdateDiaryReq struct {
	g.Meta   `path:"/diary/:id" method:"put" tags:"日记" summary:"更新日记"`
	Id       uint64   `json:"id" v:"required#请提供日记ID"`
	Title    string   `json:"title" v:"required#请输入标题"`
	Content  string   `json:"content" v:"required#请输入内容"`
	IsShared int8     `json:"isShared"`
	Mood     string   `json:"mood"`
	Weather  string   `json:"weather"`
	MediaIds []uint64 `json:"mediaIds"`
}

type UpdateDiaryRes struct {
	Success bool `json:"success"`
}

// DeleteDiaryReq 删除日记请求
type DeleteDiaryReq struct {
	g.Meta `path:"/diary/:id" method:"delete" tags:"日记" summary:"删除日记"`
	Id     uint64 `json:"id" v:"required#请提供日记ID"`
}

type DeleteDiaryRes struct {
	Success bool `json:"success"`
}

// DiaryUploadReq 日记图片上传请求
type DiaryUploadReq struct {
	g.Meta `path:"/diary/upload" method:"post" tags:"日记" summary:"上传日记图片"`
}

type DiaryUploadRes struct {
	Url      string `json:"url"`
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
}

// MediaItem 媒体附件
type MediaItem struct {
	Id        uint64 `json:"id"`
	MediaUrl  string `json:"mediaUrl"`
	MediaType string `json:"mediaType"`
}
