package user

import (
	"context"
	"fmt"

	"mallchat-go/app/user/internal/svc"
	"mallchat-go/app/user/internal/types"
	"mallchat-go/app/user/internal/utils"
	"mallchat-go/app/user/model"

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
	// 参数验证
	if len(req.Mobile) != 11 {
		return fmt.Errorf("手机号格式错误")
	}
	if len(req.NewPassword) < 6 {
		return fmt.Errorf("密码长度不能小于6位")
	}

	// 验证码校验
	code, err := l.svcCtx.SmsService.GetVerifyCode(l.ctx, req.Mobile, "reset")
	if err != nil || code != req.Code {
		return fmt.Errorf("验证码错误")
	}

	// 查找用户
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err == model.ErrNotFound {
		return fmt.Errorf("用户不存在")
	}
	if err != nil {
		return err
	}

	// 更新密码
	user.Password = utils.EncryptPassword(req.NewPassword)
	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return err
	}

	return nil
}
