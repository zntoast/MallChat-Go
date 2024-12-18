package user

import (
	"context"
	"fmt"
	"math/rand"

	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"mallchat-go/app/internal/common/errors"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送验证码
func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSmsLogic) SendSms(req *types.SendSmsReq) (resp *types.SendSmsResp, err error) {
	// 验证手机号格式
	if len(req.Mobile) != 11 {
		return nil, errors.NewValidationError("无效的手机号")
	}

	// 验证场景值
	if req.Scene != "register" && req.Scene != "reset" {
		return nil, errors.NewValidationError("无效的场景值")
	}

	// 生成6位随机验证码
	code := fmt.Sprintf("%06d", rand.Intn(1000000))

	// 存储验证码到Redis
	err = l.svcCtx.Redis.SetVerifyCode(l.ctx, req.Mobile, req.Scene, code)
	if err != nil {
		logx.Errorf("存储验证码失败: %v", err)
		return nil, errors.New(errors.UnknownError, "发送验证码失败")
	}

	return &types.SendSmsResp{
		Code: code,
	}, nil
}
