package login

import (
	"context"
	"time"

	"mallchat-go/app/user/internal/svc"
	"mallchat-go/app/user/internal/types"
	"mallchat-go/app/user/internal/utils"
	"mallchat-go/app/user/model"

	"github.com/Masterminds/squirrel"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	// 查询用户
	builder := l.svcCtx.UserModel.RowBuilder().Where(squirrel.Eq{
		"mobile":   req.Mobile,
		"password": utils.EncryptPassword(req.Password),
	})

	user, err := l.svcCtx.UserModel.FindOneByQuery(l.ctx, nil, builder)
	if err == model.ErrNotFound {
		return nil, errors.New("用户名或密码错误")
	}
	if err != nil {
		return nil, err
	}

	// 生成token
	token, expireAt, err := utils.GenerateToken(
		l.svcCtx.Config.Auth.AccessSecret,
		int64(user.Id),
		time.Duration(l.svcCtx.Config.Auth.AccessExpire)*time.Second,
	)
	if err != nil {
		return nil, err
	}

	return &types.UserLoginResp{
		AccessToken:  token,
		AccessExpire: expireAt,
		RefreshAfter: expireAt - 1200, // 提前20分钟刷新
		UserInfo: types.User{
			Id:         int64(user.Id),
			Mobile:     user.Mobile,
			Nickname:   user.Nickname.String,
			Avatar:     user.Avatar.String,
			CreateTime: user.CreateTime,
		},
	}, nil
}
