package model

import (
	"context"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ MessageModel = (*customMessageModel)(nil)

type (
	// MessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customMessageModel.
	MessageModel interface {
		messageModel
		List(ctx context.Context, receiverId uint64, lastMessageId int64, size int) ([]Message, bool, error)
	}

	customMessageModel struct {
		*defaultMessageModel
	}
)

// NewMessageModel returns a model for the database table.
func NewMessageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) MessageModel {
	return &customMessageModel{
		defaultMessageModel: newMessageModel(conn, c, opts...),
	}
}

func (m *customMessageModel) List(ctx context.Context, receiverId uint64, lastMessageId int64, size int) ([]Message, bool, error) {
	var messages []Message
	query := `SELECT * FROM message WHERE receiver_id = ? AND id < ? ORDER BY id DESC LIMIT ?`
	err := m.CachedConn.QueryRowsNoCacheCtx(ctx, &messages, query, receiverId, lastMessageId, size)
	if err != nil {
		return nil, false, err
	}

	hasMore := len(messages) == size
	return messages, hasMore, nil
}
