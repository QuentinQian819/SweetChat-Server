package promise

import (
	"chatBox/api/v1"
	"context"
	"encoding/json"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
)

type IPromiseLogic interface {
	Create(ctx context.Context, userId uint64, in *v1.CreatePromiseReq) (*v1.CreatePromiseRes, error)
	List(ctx context.Context, userId uint64, in *v1.PromiseListReq) (*v1.PromiseListRes, error)
	Get(ctx context.Context, userId uint64, promiseId uint64) (*v1.GetPromiseRes, error)
	Update(ctx context.Context, userId uint64, in *v1.UpdatePromiseReq) (*v1.UpdatePromiseRes, error)
	Delete(ctx context.Context, userId uint64, promiseId uint64) (*v1.DeletePromiseRes, error)
	ToggleComplete(ctx context.Context, userId uint64, promiseId uint64) (*v1.ToggleCompletePromiseRes, error)
}

type promiseLogicImpl struct{}

func Promise() IPromiseLogic {
	return &promiseLogicImpl{}
}

// getCoupleId 获取用户的情侣ID，如果没有绑定则返回0
func (s *promiseLogicImpl) getCoupleId(ctx context.Context, userId uint64) (uint64, error) {
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
		// 未绑定情侣关系，返回0而不是错误
		return 0, nil
	}

	return couple["id"].Uint64(), nil
}

// parseMessageIds 解析消息ID JSON字符串
func (s *promiseLogicImpl) parseMessageIds(messageIdsStr string) []uint64 {
	if messageIdsStr == "" {
		return []uint64{}
	}

	var ids []uint64
	if err := json.Unmarshal([]byte(messageIdsStr), &ids); err != nil {
		return []uint64{}
	}
	return ids
}

// formatMessageIds 格式化消息ID为JSON字符串
func (s *promiseLogicImpl) formatMessageIds(messageIds []uint64) string {
	if len(messageIds) == 0 {
		return "[]"
	}

	data, err := json.Marshal(messageIds)
	if err != nil {
		return "[]"
	}
	return string(data)
}

// Create 创建承诺
func (s *promiseLogicImpl) Create(ctx context.Context, userId uint64, in *v1.CreatePromiseReq) (*v1.CreatePromiseRes, error) {
	db := g.DB()

	coupleId, _ := s.getCoupleId(ctx, userId)
	// 即使没有绑定情侣关系也允许创建，couple_id为0

	now := gtime.Now()

	// 插入承诺
	promiseId, err := db.Model("promises").Ctx(ctx).Data(g.Map{
		"couple_id":    coupleId,
		"creator_id":   userId,
		"title":        in.Title,
		"message_ids":  s.formatMessageIds(in.MessageIds),
		"color_tag":    in.ColorTag,
		"is_completed": false,
		"created_at":   now,
		"updated_at":   now,
	}).InsertAndGetId()

	if err != nil {
		return nil, err
	}

	return &v1.CreatePromiseRes{
		PromiseId: gconv.Uint64(promiseId),
	}, nil
}

// List 获取承诺列表
func (s *promiseLogicImpl) List(ctx context.Context, userId uint64, in *v1.PromiseListReq) (*v1.PromiseListRes, error) {
	db := g.DB()

	coupleId, _ := s.getCoupleId(ctx, userId)
	// 即使没有绑定情侣关系也能查看自己创建的承诺

	// 默认分页参数
	page := in.Page
	if page <= 0 {
		page = 1
	}
	pageSize := in.PageSize
	if pageSize <= 0 || pageSize > 100 {
		pageSize = 20
	}

	// 构建查询条件：查看自己创建的承诺，且couple_id匹配（如果有情侣关系）
	query := db.Model("promises").Ctx(ctx).Where("creator_id", userId)
	if coupleId > 0 {
		query = query.Where("couple_id", coupleId)
	} else {
		query = query.Where("couple_id", 0)
	}

	// 查询总数
	total, err := query.Count()

	if err != nil {
		return nil, err
	}

	// 查询列表
	promises, err := query.
		OrderDesc("created_at").
		Page(int(page), int(pageSize)).
		All()

	if err != nil {
		return nil, err
	}

	// 转换为返回格式
	list := make([]*v1.PromiseItem, 0, len(promises))
	for _, p := range promises {
		var completedAt *time.Time
		if !p["completed_at"].IsNil() {
			t := p["completed_at"].Time()
			completedAt = &t
		}

		list = append(list, &v1.PromiseItem{
			Id:          p["id"].Uint64(),
			Title:       p["title"].String(),
			MessageIds:  s.parseMessageIds(p["message_ids"].String()),
			ColorTag:    p["color_tag"].Int(),
			IsCompleted: p["is_completed"].Bool(),
			CompletedAt: completedAt,
			CreatedAt:   p["created_at"].Time(),
			UpdatedAt:   p["updated_at"].Time(),
		})
	}

	return &v1.PromiseListRes{
		List:     list,
		Total:    int64(total),
		Page:     page,
		PageSize: pageSize,
	}, nil
}

// Get 获取承诺详情
func (s *promiseLogicImpl) Get(ctx context.Context, userId uint64, promiseId uint64) (*v1.GetPromiseRes, error) {
	db := g.DB()

	promise, err := db.Model("promises").Ctx(ctx).
		Where("id", promiseId).
		One()

	if err != nil {
		return nil, err
	}

	if promise.IsEmpty() {
		return nil, gerror.New("承诺不存在")
	}

	// 检查权限：创建者或者情侣都可以查看
	creatorId := promise["creator_id"].Uint64()
	coupleId := promise["couple_id"].Uint64()
	userCoupleId, _ := s.getCoupleId(ctx, userId)

	if creatorId != userId && coupleId != userCoupleId {
		return nil, gerror.New("无权查看此承诺")
	}

	var completedAt *time.Time
	if !promise["completed_at"].IsNil() {
		t := promise["completed_at"].Time()
		completedAt = &t
	}

	return &v1.GetPromiseRes{
		PromiseItem: v1.PromiseItem{
			Id:          promise["id"].Uint64(),
			Title:       promise["title"].String(),
			MessageIds:  s.parseMessageIds(promise["message_ids"].String()),
			ColorTag:    promise["color_tag"].Int(),
			IsCompleted: promise["is_completed"].Bool(),
			CompletedAt: completedAt,
			CreatedAt:   promise["created_at"].Time(),
			UpdatedAt:   promise["updated_at"].Time(),
		},
	}, nil
}

// Update 更新承诺
func (s *promiseLogicImpl) Update(ctx context.Context, userId uint64, in *v1.UpdatePromiseReq) (*v1.UpdatePromiseRes, error) {
	db := g.DB()

	// 检查承诺是否存在且属于当前用户
	promise, err := db.Model("promises").Ctx(ctx).
		Where("id", in.Id).
		One()

	if err != nil {
		return nil, err
	}

	if promise.IsEmpty() {
		return nil, gerror.New("承诺不存在")
	}

	if promise["creator_id"].Uint64() != userId {
		return nil, gerror.New("无权修改此承诺")
	}

	// 更新
	_, err = db.Model("promises").Ctx(ctx).
		Where("id", in.Id).
		Update(g.Map{
			"title":       in.Title,
			"message_ids": s.formatMessageIds(in.MessageIds),
			"color_tag":   in.ColorTag,
			"updated_at":  gtime.Now(),
		})

	if err != nil {
		return nil, err
	}

	return &v1.UpdatePromiseRes{
		Success: true,
	}, nil
}

// Delete 删除承诺
func (s *promiseLogicImpl) Delete(ctx context.Context, userId uint64, promiseId uint64) (*v1.DeletePromiseRes, error) {
	db := g.DB()

	// 检查承诺是否存在且属于当前用户
	promise, err := db.Model("promises").Ctx(ctx).
		Where("id", promiseId).
		One()

	if err != nil {
		return nil, err
	}

	if promise.IsEmpty() {
		return nil, gerror.New("承诺不存在")
	}

	if promise["creator_id"].Uint64() != userId {
		return nil, gerror.New("无权删除此承诺")
	}

	// 删除
	_, err = db.Model("promises").Ctx(ctx).
		Where("id", promiseId).
		Delete()

	if err != nil {
		return nil, err
	}

	return &v1.DeletePromiseRes{
		Success: true,
	}, nil
}

// ToggleComplete 切换完成状态
func (s *promiseLogicImpl) ToggleComplete(ctx context.Context, userId uint64, promiseId uint64) (*v1.ToggleCompletePromiseRes, error) {
	db := g.DB()

	// 检查承诺是否存在
	promise, err := db.Model("promises").Ctx(ctx).
		Where("id", promiseId).
		One()

	if err != nil {
		return nil, err
	}

	if promise.IsEmpty() {
		return nil, gerror.New("承诺不存在")
	}

	// 检查权限：创建者或者情侣都可以切换完成状态
	creatorId := promise["creator_id"].Uint64()
	coupleId := promise["couple_id"].Uint64()
	userCoupleId, _ := s.getCoupleId(ctx, userId)

	if creatorId != userId && coupleId != userCoupleId && coupleId > 0 {
		return nil, gerror.New("无权操作此承诺")
	}

	// 获取当前状态
	currentStatus := promise["is_completed"].Bool()
	newStatus := !currentStatus

	// 准备更新数据
	updateData := g.Map{
		"is_completed": newStatus,
		"updated_at":   gtime.Now(),
	}

	// 如果标记为完成，设置完成时间；如果取消完成，清空完成时间
	if newStatus {
		updateData["completed_at"] = gtime.Now()
	} else {
		updateData["completed_at"] = nil
	}

	// 更新
	_, err = db.Model("promises").Ctx(ctx).
		Where("id", promiseId).
		Update(updateData)

	if err != nil {
		return nil, err
	}

	// 获取更新后的数据
	var completedAt *time.Time
	if newStatus {
		now := time.Now()
		completedAt = &now
	}

	return &v1.ToggleCompletePromiseRes{
		Success:     true,
		IsCompleted: newStatus,
		CompletedAt: completedAt,
	}, nil
}
