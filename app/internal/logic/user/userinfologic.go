package user

import (
	"context"
	"net/http"

	"mallchat-go/app/internal/ercode"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
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
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(req.UserId))
	if err != nil {
		return nil, ercode.New(http.StatusInternalServerError, "获取用户信息失败", zap.Error(err))
	}
	if user == nil {
		return nil, ercode.New(http.StatusNotFound, "用户不存在")
	}
	resp = new(types.UserInfoResp)
	copier.Copy(resp, user)
	return
}
