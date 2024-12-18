package user

import (
	"context"

	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
