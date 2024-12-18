package login

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"mallchat-go/app/internal/common/errors"
	"mallchat-go/app/internal/model"
	"mallchat-go/app/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// register
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRegisterReq) (resp *types.UserRegisterResp, err error) {
	// 验证手机号格式
	if len(req.Mobile) != 11 {
		return nil, errors.NewValidationError("无效的手机号")
	}

	// 验证密码长度
	if len(req.Password) < 6 || len(req.Password) > 20 {
		return nil, errors.NewValidationError("密码长度应在6-20个字符之间")
	}

	// 验证验证码
	code, err := l.svcCtx.Redis.GetVerifyCode(l.ctx, req.Mobile, "register")
	if err != nil || code != req.Code {
		return nil, errors.NewValidationError("验证码错误或已过期")
	}

	// 检查手机号是否已注册
	exists, err := l.svcCtx.UserModel.ExistsByMobile(l.ctx, req.Mobile)
	if err != nil {
		return nil, err
	}
	if exists {
		return nil, errors.NewValidationError("该手机号已注册")
	}

	// 创建用户
	hashedPassword := utils.HashPassword(req.Password)
	user := &model.User{
		Mobile:     req.Mobile,
		Password:   hashedPassword,
		Nickname:   sql.NullString{String: fmt.Sprintf("用户%s", req.Mobile[7:]), Valid: true},
		CreateTime: time.Now().Unix(),
	}
	result, err := l.svcCtx.UserModel.Insert(l.ctx, user)
	if err != nil {
		return nil, err
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// 生成token
	token, expireTime, refreshTime, err := l.svcCtx.JWT.GenerateToken(userId)
	if err != nil {
		return nil, errors.New(errors.UnknownError, "生成token失败")
	}

	// 存储token到Redis
	err = l.svcCtx.Redis.SetUserToken(l.ctx, userId, token)
	if err != nil {
		logx.Errorf("存储token失败: %v", err)
	}

	return &types.UserRegisterResp{
		AccessToken:  token,
		AccessExpire: expireTime,
		RefreshAfter: refreshTime,
	}, nil
}
