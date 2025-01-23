package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BlacksModel = (*customBlacksModel)(nil)

type (
	// BlacksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBlacksModel.
	BlacksModel interface {
		blacksModel
	}

	customBlacksModel struct {
		*defaultBlacksModel
	}
)

// NewBlacksModel returns a model for the database table.
func NewBlacksModel(conn sqlx.SqlConn) BlacksModel {
	return &customBlacksModel{
		defaultBlacksModel: newBlacksModel(conn),
	}
}
