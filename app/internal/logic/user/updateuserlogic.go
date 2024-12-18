package user

import (
	"context"
	"database/sql"
	"fmt"

	"mallchat-go/app/internal/common"
	"mallchat-go/app/internal/model"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新用户信息
func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserLogic) UpdateUser(req *types.UpdateUserReq) error {
	// 从上下文获取用户ID
	userId, ok := common.GetUserIDFromContext(l.ctx)
	if !ok {
		return fmt.Errorf("无效的用户ID")
	}

	// 验证昵称长度
	if len(req.Nickname) > 0 && (len(req.Nickname) < 2 || len(req.Nickname) > 20) {
		return fmt.Errorf("昵称长度应在2-20个字符之间")
	}

	// 验证头像URL格式
	if len(req.Avatar) > 0 && len(req.Avatar) > 255 {
		return fmt.Errorf("头像URL过长")
	}

	// 更新数据库中的用户信息
	user := &model.User{
		Id:       uint64(userId),
		Nickname: sql.NullString{String: req.Nickname, Valid: true},
		Avatar:   sql.NullString{String: req.Avatar, Valid: true},
	}
	err := l.svcCtx.UserModel.Update(l.ctx, user)
	if err != nil {
		return err
	}

	return nil
}
