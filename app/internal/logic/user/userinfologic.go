package user

import (
	"context"
	"net/http"

	"mallchat-go/app/internal/ercode"
	modelUser "mallchat-go/app/internal/model/user"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户详情
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo(req *types.UserInfoReq) (resp *types.UserInfoResp, err error) {
	user := modelUser.User{}
	err = l.svcCtx.Db.Where("id = ?", req.UserId).First(&user).Error
	if err != nil {
		return nil, ercode.New(http.StatusNotFound, err.Error())
	}
	resp = new(types.UserInfoResp)
	copier.Copy(resp, user)
	return
}
