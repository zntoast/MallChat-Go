package user

import (
	"context"
	"time"

	"mallchat-go/app/internal/errors"
	"mallchat-go/app/internal/middleware"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
)

type ModifyNameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 修改用户名
func NewModifyNameLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ModifyNameLogic {
	return &ModifyNameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ModifyNameLogic) ModifyName(req *types.ModifyNameReq) error {
	newName := req.Name
	userId := middleware.GetAuthRespFromCtx(l.ctx)

	// // 检验用户名是否合法
	has, _ := l.svcCtx.Filter.FindIn(newName)
	if has {
		return errors.New(errors.ErrStatusBadRequest, "包含违禁词，请修改用户名~~")
	}

	query := l.svcCtx.UserModel.SelectBuilder()
	query.Where(
		squirrel.And{
			squirrel.Eq{"name": newName},
			squirrel.NotEq{"id": userId},
		},
	)
	// 检验用户名是否存在
	count, err := l.svcCtx.UserModel.FindCount(l.ctx, query, "id")
	if err != nil {
		return errors.New(errors.SysDBError, "系统内部错误", zap.Error(err))
	}

	if count > 0 {
		return errors.New(errors.ErrStatusBadRequest, "用户名已存在，请修改用户名~~")
	}

	user, err := l.svcCtx.UserModel.FindOne(l.ctx, userId)
	if err != nil {
		return errors.New(errors.SysDBError, "查询用户信息失败~", zap.Error(err))
	}
	if user == nil {
		return errors.New(errors.ErrDataNotFound, "account is not exist", zap.Any("user_id", userId))
	}

	user.Name = newName
	user.UpdatedAt = time.Now()
	_, err = l.svcCtx.UserModel.Update(l.ctx, nil, user)
	if err != nil {
		return errors.New(errors.SysDBError, "更新用户名失败~", zap.Error(err))
	}
	return nil
}
