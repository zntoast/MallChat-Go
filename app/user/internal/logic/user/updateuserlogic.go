package user

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"mallchat-go/app/user/internal/svc"
	"mallchat-go/app/user/internal/types"
	"mallchat-go/app/user/model"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) error {
	// 从context中获取用户ID
	userIdStr := l.ctx.Value("X-User-ID").(string)
	userId, _ := strconv.ParseUint(userIdStr, 10, 64)

	// 查找用户
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err == model.ErrNotFound {
		return fmt.Errorf("用户不存在")
	}
	if err != nil {
		return err
	}

	// 更新用户信息
	if req.Nickname != "" {
		user.Nickname = sql.NullString{String: req.Nickname, Valid: true}
	}
	if req.Avatar != "" {
		user.Avatar = sql.NullString{String: req.Avatar, Valid: true}
	}
	user.UpdateTime = time.Now().Unix()

	err = l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return err
	}

	return nil
}
