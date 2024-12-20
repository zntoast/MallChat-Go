package group

import (
	"context"

	"mallchat-go/app/internal/model"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type JoinGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 加入群组
func NewJoinGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *JoinGroupLogic {
	return &JoinGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *JoinGroupLogic) JoinGroup(req *types.JoinGroupReq) (resp *types.JoinGroupResp, err error) {
	// 检查用户是否已经在群组中
	member, err := l.svcCtx.GroupModel.FindOneByGroupIdUserId(l.ctx, req.GroupId, req.UserId)
	if err != nil && err != model.ErrNotFound {
		return nil, err // 返回错误
	}
	if member != nil {
		return &types.JoinGroupResp{
			Message: "用户已在群组中",
		}, nil
	}

	// 如果用户不在群组中，执行加入操作
	err = l.svcCtx.GroupModel.AddMember(req.GroupId, req.UserId)
	if err != nil {
		return nil, err // 返回错误
	}

	// 返回成功消息
	resp = &types.JoinGroupResp{
		Message: "成功加入群组",
	}
	return
}
