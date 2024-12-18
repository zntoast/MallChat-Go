package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ OfflineMessageModel = (*customOfflineMessageModel)(nil)

type (
	// OfflineMessageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customOfflineMessageModel.
	OfflineMessageModel interface {
		offlineMessageModel
	}

	customOfflineMessageModel struct {
		*defaultOfflineMessageModel
	}
)

// NewOfflineMessageModel returns a model for the database table.
func NewOfflineMessageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) OfflineMessageModel {
	return &customOfflineMessageModel{
		defaultOfflineMessageModel: newOfflineMessageModel(conn, c, opts...),
	}
}
