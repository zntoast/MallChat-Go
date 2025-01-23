package user

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"mallchat-go/app/internal/svc"
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
	// todo: add your logic here and delete this line

	return nil
}
