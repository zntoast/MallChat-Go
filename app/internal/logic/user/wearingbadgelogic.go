package user

import (
	"context"

	"mallchat-go/app/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type WearingBadgeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 佩戴徽章
func NewWearingBadgeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WearingBadgeLogic {
	return &WearingBadgeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WearingBadgeLogic) WearingBadge() error {

	return nil
}
