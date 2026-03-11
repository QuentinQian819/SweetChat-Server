package diary

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
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
)

type IDiaryLogic interface {
	Create(ctx context.Context, userId uint64, in *v1.CreateDiaryReq) (*v1.CreateDiaryRes, error)
	List(ctx context.Context, userId uint64, in *v1.DiaryListReq) (*v1.DiaryListRes, error)
	Get(ctx context.Context, userId uint64, diaryId uint64) (*v1.GetDiaryRes, error)
	Update(ctx context.Context, userId uint64, in *v1.UpdateDiaryReq) (*v1.UpdateDiaryRes, error)
	Delete(ctx context.Context, userId uint64, diaryId uint64) (*v1.DeleteDiaryRes, error)
	Upload(ctx context.Context, userId uint64, file *ghttp.UploadFile) (*v1.DiaryUploadRes, error)
}

type diaryLogicImpl struct{}

func Diary() IDiaryLogic {
	return &diaryLogicImpl{}
}

// getCoupleId 获取用户的情侣ID
func (s *diaryLogicImpl) getCoupleId(ctx context.Context, userId uint64) (uint64, error) {
	db := g.DB()

	couple, err := db.Model("couples").Ctx(ctx).
		Where("user1_id", userId).
		WhereOr("user2_id", userId).
		Where("status", 1).
		One()

	if err != nil {
		return 0, err
	}

	if couple.IsEmpty() {
		return 0, gerror.New("未绑定情侣关系")
	}

	return couple["id"].Uint64(), nil
}

// Create 创建日记
func (s *diaryLogicImpl) Create(ctx context.Context, userId uint64, in *v1.CreateDiaryReq) (*v1.CreateDiaryRes, error) {
	db := g.DB()

	coupleId, err := s.getCoupleId(ctx, userId)
	if err != nil {
		return nil, err
	}

	now := gtime.Now()

	// 插入日记
	diaryId, err := db.Model("diaries").Ctx(ctx).Data(g.Map{
		"couple_id":  coupleId,
		"author_id":  userId,
		"title":      in.Title,
		"content":    in.Content,
		"is_shared":  in.IsShared,
		"mood":       in.Mood,
		"weather":    in.Weather,
		"created_at": now,
		"updated_at": now,
	}).InsertAndGetId()

	if err != nil {
		return nil, err
	}

	return &v1.CreateDiaryRes{
		DiaryId: gconv.Uint64(diaryId),
	}, nil
}

// List 获取日记列表
func (s *diaryLogicImpl) List(ctx context.Context, userId uint64, in *v1.DiaryListReq) (*v1.DiaryListRes, error) {
	db := g.DB()

	coupleId, err := s.getCoupleId(ctx, userId)
	if err != nil {
		return nil, err
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

	// 查询总数
	total, err := db.Model("diaries").Ctx(ctx).
		Where("couple_id", coupleId).
		Where("author_id", userId).
		Count()

	if err != nil {
		return nil, err
	}

	// 查询列表
	diaries, err := db.Model("diaries").Ctx(ctx).
		Where("couple_id", coupleId).
		WhereOr("author_id", userId).
		Where("is_shared", 1).
		OrderDesc("created_at").
		Page(int(page), int(pageSize)).
		All()

	if err != nil {
		return nil, err
	}

	// 转换为返回格式
	list := make([]*v1.DiaryItem, 0, len(diaries))
	for _, d := range diaries {
		list = append(list, &v1.DiaryItem{
			Id:        d["id"].Uint64(),
			Title:     d["title"].String(),
			Content:   d["content"].String(),
			IsShared:  d["is_shared"].Int() == 1,
			Mood:      d["mood"].String(),
			Weather:   d["weather"].String(),
			CreatedAt: d["created_at"].Time(),
			UpdatedAt: d["updated_at"].Time(),
			Media:     []*v1.MediaItem{},
		})
	}

	return &v1.DiaryListRes{
		List:     list,
		Total:    int64(total),
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// Get 获取日记详情
func (s *diaryLogicImpl) Get(ctx context.Context, userId uint64, diaryId uint64) (*v1.GetDiaryRes, error) {
	db := g.DB()

	diary, err := db.Model("diaries").Ctx(ctx).
		Where("id", diaryId).
		One()

	if err != nil {
		return nil, err
	}

	if diary.IsEmpty() {
		return nil, gerror.New("日记不存在")
	}

	// 检查权限
	authorId := diary["author_id"].Uint64()
	isShared := diary["is_shared"].Int() == 1

	if authorId != userId && !isShared {
		return nil, gerror.New("无权查看此日记")
	}

	return &v1.GetDiaryRes{
		DiaryItem: v1.DiaryItem{
			Id:        diary["id"].Uint64(),
			Title:     diary["title"].String(),
			Content:   diary["content"].String(),
			IsShared:  isShared,
			Mood:      diary["mood"].String(),
			Weather:   diary["weather"].String(),
			CreatedAt: diary["created_at"].Time(),
			UpdatedAt: diary["updated_at"].Time(),
			Media:     []*v1.MediaItem{},
		},
	}, nil
}

// Update 更新日记
func (s *diaryLogicImpl) Update(ctx context.Context, userId uint64, in *v1.UpdateDiaryReq) (*v1.UpdateDiaryRes, error) {
	db := g.DB()

	// 检查日记是否存在且属于当前用户
	diary, err := db.Model("diaries").Ctx(ctx).
		Where("id", in.Id).
		One()

	if err != nil {
		return nil, err
	}

	if diary.IsEmpty() {
		return nil, gerror.New("日记不存在")
	}

	if diary["author_id"].Uint64() != userId {
		return nil, gerror.New("无权修改此日记")
	}

	// 更新
	_, err = db.Model("diaries").Ctx(ctx).
		Where("id", in.Id).
		Update(g.Map{
			"title":      in.Title,
			"content":    in.Content,
			"is_shared":  in.IsShared,
			"mood":       in.Mood,
			"weather":    in.Weather,
			"updated_at": gtime.Now(),
		})

	if err != nil {
		return nil, err
	}

	return &v1.UpdateDiaryRes{
		Success: true,
	}, nil
}

// Delete 删除日记
func (s *diaryLogicImpl) Delete(ctx context.Context, userId uint64, diaryId uint64) (*v1.DeleteDiaryRes, error) {
	db := g.DB()

	// 检查日记是否存在且属于当前用户
	diary, err := db.Model("diaries").Ctx(ctx).
		Where("id", diaryId).
		One()

	if err != nil {
		return nil, err
	}

	if diary.IsEmpty() {
		return nil, gerror.New("日记不存在")
	}

	if diary["author_id"].Uint64() != userId {
		return nil, gerror.New("无权删除此日记")
	}

	// 删除
	_, err = db.Model("diaries").Ctx(ctx).
		Where("id", diaryId).
		Delete()

	if err != nil {
		return nil, err
	}

	return &v1.DeleteDiaryRes{
		Success: true,
	}, nil
}

// Upload 上传日记图片
func (s *diaryLogicImpl) Upload(ctx context.Context, userId uint64, file *ghttp.UploadFile) (*v1.DiaryUploadRes, error) {
	if file == nil {
		return nil, gerror.New("请选择文件")
	}

	// 检查文件大小
	config := g.Cfg()
	maxSize := config.MustGet(ctx, "upload.max_size").Int64()
	if file.Size > maxSize {
		return nil, gerror.Newf("文件大小不能超过%dMB", maxSize/1024/1024)
	}

	// 检查文件类型（只允许图片）
	allowedTypes := config.MustGet(ctx, "upload.allowed_image_types").Strings()
	contentType := file.Header.Get("Content-Type")

	isAllowed := false
	for _, t := range allowedTypes {
		if strings.HasPrefix(contentType, t) {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		return nil, gerror.New("只支持图片格式")
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

	return &v1.DiaryUploadRes{
		Url:      url,
		Filename: file.Filename,
		Size:     file.Size,
	}, nil
}
