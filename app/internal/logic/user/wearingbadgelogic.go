package user

import (
	"context"
	"time"

	"mallchat-go/app/internal/errors"
	"mallchat-go/app/internal/middleware"
	"mallchat-go/app/internal/model"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

type WearingBadgeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWearingBadgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WearingBadgeLogic {
	return &WearingBadgeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WearingBadgeLogic) WearingBadge(req *types.WearingBadgeReq) error {

	user := middleware.GetAuthRespFromCtx(l.ctx)
	query := l.svcCtx.UserBackpacksDb.SelectBuilder()

	query.Where(
		squirrel.And{
			squirrel.Eq{"item_id": req.BadgeId},
			squirrel.Eq{"uid": user.Id},
			squirrel.Eq{"status": model.UserBackpacksStatusValid},
		},
	)
	userBackpacks, err := l.svcCtx.UserBackpacksDb.FindAll(l.ctx, query, "id")
	if err != nil {
		return errors.New(errors.SysDBError, "系统内部错误", zap.Error(err))
	}
	if len(userBackpacks) == 0 {
		return errors.New(errors.ErrDataNotFound, "该用户未拥有该徽章", zap.Int64("badge_id", req.BadgeId))
	}

	userBackpack := userBackpacks[0]

	item, err := l.svcCtx.ItemConfigsDb.FindOne(l.ctx, uint64(userBackpack.ItemId))
	if err != nil {
		return errors.New(errors.SysDBError, "系统内部错误", zap.Error(err))
	}

	if item == nil || item.Type != model.ItemConfigsTypeBadge {
		return errors.New(errors.ErrDataNotFound, "该徽章不可佩戴", zap.Int64("item_id", int64(item.Id)))
	}

	user.ItemId = int64(item.Id)
	user.UpdatedAt = time.Now()
	_, err = l.svcCtx.UserDb.Update(l.ctx, nil, user)
	if err != nil {
		return errors.New(errors.SysDBError, "徽章佩戴失败", zap.Error(err))
	}

	return nil
}
