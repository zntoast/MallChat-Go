package message

import (
	"context"
	"fmt"

	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取消息列表
func NewGetMessageListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageListLogic {
	return &GetMessageListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetMessageListLogic) GetMessageList(req *types.GetMessageListReq) (resp *types.GetMessageListResp, err error) {
	// 验证参数
	if req.ReceiverId <= 0 {
		return nil, fmt.Errorf("无效的接收者ID")
	}

	if req.Size <= 0 || req.Size > 100 {
		req.Size = 20 // 使用默认值
	}

	// 从数据库查询消息列表
	messages, hasMore, err := l.svcCtx.MessageModel.List(l.ctx, uint64(req.ReceiverId), req.LastMessageId, int(req.Size))
	if err != nil {
		return nil, err
	}

	// 转换消息类型
	var responseMessages []types.Message
	for _, msg := range messages {
		responseMessages = append(responseMessages, types.Message{
			Id:         int64(msg.Id),
			SenderId:   int64(msg.SenderId),
			ReceiverId: int64(msg.ReceiverId),
			Content:    msg.Content,
			Type:       int32(msg.Type),
			CreateTime: msg.CreateTime,
		})
	}

	return &types.GetMessageListResp{
		Messages: responseMessages,
		HasMore:  hasMore,
	}, nil
}
