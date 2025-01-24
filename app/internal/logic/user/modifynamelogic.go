package user

import (
	"context"
	"net/http"

	"mallchat-go/app/internal/ercode"
	"mallchat-go/app/internal/middleware"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/logx"
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
		return ercode.New(http.StatusBadRequest, "包含违禁词，请修改用户名~~")
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
		return ercode.NewSysError("")
	}

	if count > 0 {
		return ercode.New(http.StatusBadRequest, "用户名已存在，请修改用户名~~")
	}

	// l.svcCtx.UserModel.

	// err = l.svcCtx.Db.Model(modelUser.User{}).Where("id = ?", userId).Update("name", newName).Error
	// if err != nil {
	// 	return ercode.New(http.StatusInternalServerError, "系统错误，请稍后再试~~", zap.Error(err))
	// }
	return nil
}
