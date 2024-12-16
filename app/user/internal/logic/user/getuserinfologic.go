package user

import (
	"context"
	"fmt"
	"strconv"

	"mallchat-go/app/user/internal/svc"
	"mallchat-go/app/user/internal/types"
	"mallchat-go/app/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo() (resp *types.GetUserInfoResp, err error) {
	// 从context中获取用户ID
	userIdStr := l.ctx.Value("X-User-ID").(string)
	userId, _ := strconv.ParseUint(userIdStr, 10, 64)

	// 查找用户
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err == model.ErrNotFound {
		return nil, fmt.Errorf("用户不存在")
	}
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
