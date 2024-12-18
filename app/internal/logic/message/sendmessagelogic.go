package message

import (
	"context"
	"fmt"
	"time"

	"mallchat-go/app/internal/model"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"
	wsTypes "mallchat-go/app/internal/types/ws"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送消息
func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendMessageLogic) SendMessage(req *types.SendMessageReq) (resp *types.SendMessageResp, err error) {
	// 验证接收者ID
	if req.ReceiverId <= 0 {
		return nil, fmt.Errorf("无效的接收者ID")
	}

	// 保存消息到数据库
	message := &model.Message{
		SenderId:   uint64(req.SenderId),
		ReceiverId: uint64(req.ReceiverId),
		Content:    req.Content,
		Type:       int64(req.Type),
		CreateTime: time.Now().Unix(),
	}
	result, err := l.svcCtx.MessageModel.Insert(l.ctx, message)
	if err != nil {
		return nil, err
	}

	messageId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	// 通过WebSocket推送消息
	wsMessage := &wsTypes.Message{
		Type:      req.Type,
		SenderId:  req.SenderId,
		Content:   req.Content,
		Timestamp: message.CreateTime,
	}
	l.svcCtx.WS.SendToUser(req.ReceiverId, wsMessage)

	return &types.SendMessageResp{
		MessageId: messageId,
	}, nil
}
