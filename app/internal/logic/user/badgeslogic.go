package user

import (
	"context"

	"mallchat-go/app/internal/errors"
	"mallchat-go/app/internal/model"
	"mallchat-go/app/internal/pkg/utils"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

type BadgesLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewBadgesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BadgesLogic {
	return &BadgesLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *BadgesLogic) Badges(req *types.BadgesReq) (resp *types.BadgesItemsResp, err error) {

	itemConfigs, err := l.svcCtx.ItemConfigsDb.FindAll(l.ctx, squirrel.SelectBuilder{}.Where(
		squirrel.And{
			squirrel.Eq{"type": model.ItemConfigsTypeBadge},
		},
	), "id")
	if err != nil {
		return nil, errors.New(errors.SysDBError, "获取徽章配置失败", zap.Error(err))
	}

	if len(itemConfigs) == 0 {
		return nil, nil
	}

	itemIds := utils.ExtractToSlice(itemConfigs, func(item *model.ItemConfigs) uint64 { return item.Id })
	userBackpacks, err := l.svcCtx.UserBackpacksDb.FindAll(l.ctx, squirrel.SelectBuilder{}.Where(
		squirrel.And{
			squirrel.Eq{"uid": req.Uid},
			squirrel.Eq{"item_id": itemIds},
		},
	), "id")
	if err != nil {
		return nil, errors.New(errors.SysDBError, "获取用户背包徽章失败", zap.Error(err))
	}

	user, err := l.svcCtx.UserDb.FindOne(l.ctx, uint64(req.Uid))
	if err != nil {
		return nil, errors.New(errors.SysDBError, "获取用户信息失败", zap.Error(err))
	}
	return buildBadgeResp(itemConfigs, userBackpacks, user), nil
}

func buildBadgeResp(itemConfigs []*model.ItemConfigs, backpacks []*model.UserBackpacks, user *model.Users) (resp *types.BadgesItemsResp) {
	if user == nil {
		return nil
	}
	resp = new(types.BadgesItemsResp)
	resp.Items = make([]types.BadgesItem, 0)

	return nil
}
