package user

import (
	"context"
	"fmt"

	"mallchat-go/app/internal/middleware"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

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
	// 检验用户名是否合法

	has, _ := l.svcCtx.Filter.FindIn(newName)
	if has {
		return fmt.Errorf("包含违禁词，请修改用户名~~")
	}

	// 检验用户名是否存在
	var count int64 = 0
	err := l.svcCtx.Db.Where("name = ? and id <>?", newName, userId).Count(&count).Error
	if err != nil {
		return fmt.Errorf("名字已经被抢占了，请换一个哦~~")
	}

	err = l.svcCtx.Db.Where("id = ?", userId).Update("name", newName).Error
	if err != nil {
		return fmt.Errorf("修改用户名失败，请稍后再试~~")
	}
	return nil
}
