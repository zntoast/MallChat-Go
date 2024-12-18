package login

import (
	"context"
	"fmt"

	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"
	"mallchat-go/app/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// login
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.UserLoginReq) (resp *types.UserLoginResp, err error) {
	// 验证手机号格式
	if len(req.Mobile) != 11 {
		return nil, fmt.Errorf("无效的手机号")
	}

	// 验证密码长度
	if len(req.Password) < 6 || len(req.Password) > 20 {
		return nil, fmt.Errorf("密码长度应在6-20个字符之间")
	}

	// TODO: 验证用户密码
	user, err := l.svcCtx.UserModel.FindByMobile(req.Mobile)
	if err != nil {
		return nil, fmt.Errorf("用户不存在")
	}
	if !utils.ValidatePassword(req.Password, user.Password) {
		return nil, fmt.Errorf("密码错误")
	}

	// 生成JWT token
	token, expireTime, refreshTime, err := l.svcCtx.JWT.GenerateToken(int64(user.Id))
	if err != nil {
		return nil, fmt.Errorf("生成token失败")
	}

	// 存储token到Redis
	err = l.svcCtx.Redis.SetUserToken(l.ctx, int64(user.Id), token)
	if err != nil {
		logx.Errorf("存储token失败: %v", err)
	}

	return &types.UserLoginResp{
		AccessToken:  token,
		AccessExpire: expireTime,
		RefreshAfter: refreshTime,
		UserInfo: types.User{
			Id:         int64(user.Id),
			Mobile:     user.Mobile,
			Nickname:   user.Nickname.String,
			Avatar:     user.Avatar.String,
			CreateTime: user.CreateTime,
		},
	}, nil
}
