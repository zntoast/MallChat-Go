package login

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"mallchat-go/app/user/internal/svc"
	"mallchat-go/app/user/internal/types"
	"mallchat-go/app/user/model"
	"mallchat-go/app/user/utils"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.UserRegisterReq) (resp *types.UserRegisterResp, err error) {
	// 添加日志
	logx.Infof("用户注册请求: mobile=%s", req.Mobile)
	defer func() {
		if err != nil {
			logx.Errorf("用户注册失败: mobile=%s, err=%v", req.Mobile, err)
		} else {
			logx.Infof("用户注册成功: mobile=%s", req.Mobile)
		}
	}()

	// 参数验证
	if len(req.Mobile) != 11 {
		return nil, errors.New("手机号格式错误")
	}
	if len(req.Password) < 6 {
		return nil, errors.New("密码长度不能小于6位")
	}

	// 验证码校验
	code, err := l.svcCtx.SmsService.GetVerifyCode(l.ctx, req.Mobile, "register")
	if err != nil || code != req.Code {
		return nil, errors.New("验证码错误")
	}

	// 检查手机号是否已注册
	exists, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, req.Mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, err
	}
	if exists != nil {
		return nil, errors.New("手机号已注册")
	}

	var userId int64
	err = l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		// 创建用户
		user := &model.User{
			Mobile:     req.Mobile,
			Password:   utils.EncryptPassword(req.Password),
			Nickname:   sql.NullString{String: fmt.Sprintf("用户%d", time.Now().Unix()), Valid: true},
			CreateTime: time.Now().Unix(),
			UpdateTime: time.Now().Unix(),
		}

		result, err := l.svcCtx.UserModel.Insert(ctx, user)
		if err != nil {
			return err
		}

		userId, err = result.LastInsertId()
		if err != nil {
			return err
		}

		// 可以在这里添加其他需要在事务中执行的操作
		return nil
	})

	if err != nil {
		return nil, err
	}

	// 生成token
	token, expireAt, err := utils.GenerateToken(
		l.svcCtx.Config.Auth.AccessSecret,
		userId,
		time.Duration(l.svcCtx.Config.Auth.AccessExpire)*time.Second,
	)
	if err != nil {
		return nil, err
	}

	return &types.UserRegisterResp{
		AccessToken:  token,
		AccessExpire: expireAt,
		RefreshAfter: expireAt - 1200, // 提前20分钟刷新
	}, nil
}
