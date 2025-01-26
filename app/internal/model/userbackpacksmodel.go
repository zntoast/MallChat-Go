package model

import (
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserBackpacksModel = (*customUserBackpacksModel)(nil)

const (
	UserBackpacksStatusValid   = 0 // 未失效
	UserBackpacksStatusInvalid = 1 // 已失效
)

type (
	// UserBackpacksModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserBackpacksModel.
	UserBackpacksModel interface {
		userBackpacksModel
	}

	customUserBackpacksModel struct {
		*defaultUserBackpacksModel
	}
)

// NewUserBackpacksModel returns a model for the database table.
func NewUserBackpacksModel(conn sqlx.SqlConn) UserBackpacksModel {
	return &customUserBackpacksModel{
		defaultUserBackpacksModel: newUserBackpacksModel(conn),
	}
}
