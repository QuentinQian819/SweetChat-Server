package diary

import (
	"chatBox/api/v1"
	"chatBox/internal/logic/diary"
	"context"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/util/gconv"
)

type ControllerV1 struct{}

func NewV1() *ControllerV1 {
	return &ControllerV1{}
}

// Create 创建日记
func (c *ControllerV1) Create(ctx context.Context, req *v1.CreateDiaryReq) (res *v1.CreateDiaryRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return diary.Diary().Create(ctx, userId, req)
}

// List 获取日记列表
func (c *ControllerV1) List(ctx context.Context, req *v1.DiaryListReq) (res *v1.DiaryListRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return diary.Diary().List(ctx, userId, req)
}

// Get 获取日记详情
func (c *ControllerV1) Get(ctx context.Context, req *v1.GetDiaryReq) (res *v1.GetDiaryRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return diary.Diary().Get(ctx, userId, req.Id)
}

// Update 更新日记
func (c *ControllerV1) Update(ctx context.Context, req *v1.UpdateDiaryReq) (res *v1.UpdateDiaryRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return diary.Diary().Update(ctx, userId, req)
}

// Delete 删除日记
func (c *ControllerV1) Delete(ctx context.Context, req *v1.DeleteDiaryReq) (res *v1.DeleteDiaryRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	return diary.Diary().Delete(ctx, userId, req.Id)
}

// Upload 上传日记图片
func (c *ControllerV1) Upload(ctx context.Context, req *v1.DiaryUploadReq) (res *v1.DiaryUploadRes, err error) {
	r := ghttp.RequestFromCtx(ctx)
	userId := gconv.Uint64(r.GetCtxVar("userId"))
	file := r.GetUploadFile("file")
	if file == nil {
		return nil, gerror.New("请选择文件")
	}
	return diary.Diary().Upload(ctx, userId, file)
}
