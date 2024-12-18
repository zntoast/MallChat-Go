package user

import (
	"context"
	"fmt"

	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"mallchat-go/app/internal/common"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取用户信息
func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResp, err error) {
	// 从上下文获取用户ID
	userId, ok := common.GetUserIDFromContext(l.ctx)
	if !ok {
		return nil, fmt.Errorf("无效的用户ID")
	}

	// TODO: 从数据库获取用户信息
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(userId))
	if err != nil {
		return nil, err
	}

	return &types.GetUserInfoResp{
		Id:         int64(user.Id),
		Mobile:     user.Mobile,
		Nickname:   user.Nickname.String,
		Avatar:     user.Avatar.String,
		CreateTime: user.CreateTime,
	}, nil
}
