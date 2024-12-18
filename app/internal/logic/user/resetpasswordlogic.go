package user

import (
	"context"
	"fmt"

	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"
	"mallchat-go/app/internal/utils"

	"github.com/zeromicro/go-zero/core/logx"
)

type ResetPasswordLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 重置密码
func NewResetPasswordLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ResetPasswordLogic {
	return &ResetPasswordLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ResetPasswordLogic) ResetPassword(req *types.ResetPasswordReq) error {
	// 验证手机号格式
	if len(req.Mobile) != 11 {
		return fmt.Errorf("无效的手机号")
	}

	// 验证密码长度
	if len(req.NewPassword) < 6 || len(req.NewPassword) > 20 {
		return fmt.Errorf("密码长度应在6-20个字符之间")
	}

	// 验证验证码
	if len(req.Code) != 6 {
		return fmt.Errorf("无效的验证码")
	}

	code, err := l.svcCtx.Redis.GetVerifyCode(l.ctx, req.Mobile, "reset")
	if err != nil || code != req.Code {
		return fmt.Errorf("验证码错误或已过期")
	}

	// 更新用户密码
	hashedPassword := utils.HashPassword(req.NewPassword)
	err = l.svcCtx.UserModel.UpdatePassword(req.Mobile, hashedPassword)
	if err != nil {
		return err
	}

	return nil
}
