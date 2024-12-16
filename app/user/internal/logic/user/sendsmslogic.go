package user

import (
	"context"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	"mallchat-go/app/user/internal/svc"
	"mallchat-go/app/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendSmsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSendSmsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendSmsLogic {
	return &SendSmsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendSmsLogic) SendSms(req *types.SendSmsReq) (resp *types.SendSmsResp, err error) {
	// 验证手机号格式
	if !regexp.MustCompile(`^1[3-9]\d{9}$`).MatchString(req.Mobile) {
		return nil, fmt.Errorf("手机号格式错误")
	}

	// 验证场景值
	if req.Scene != "register" && req.Scene != "reset" {
		return nil, fmt.Errorf("不支持的验证码场景")
	}

	// 生成验证码
	code := fmt.Sprintf("%06d", rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000))

	// 存储验证码
	err = l.svcCtx.SmsService.SetVerifyCode(l.ctx, req.Mobile, req.Scene, code)
	if err != nil {
		logx.Errorf("存储验证码失败: %v", err)
		return nil, fmt.Errorf("发送验证码失败")
	}

	// TODO: 调用短信服务发送验证码
	// 这里先返回验证码，方便测试
	return &types.SendSmsResp{
		Code: code,
	}, nil
}
