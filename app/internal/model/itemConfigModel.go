package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ ItemConfigModel = (*customItemConfigModel)(nil)

type (
	// ItemConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customItemConfigModel.
	ItemConfigModel interface {
		itemConfigModel
		withSession(session sqlx.Session) ItemConfigModel
	}

	customItemConfigModel struct {
		*defaultItemConfigModel
	}
)

// NewItemConfigModel returns a model for the database table.
func NewItemConfigModel(conn sqlx.SqlConn) ItemConfigModel {
	return &customItemConfigModel{
		defaultItemConfigModel: newItemConfigModel(conn),
	}
}

func (m *customItemConfigModel) withSession(session sqlx.Session) ItemConfigModel {
	return NewItemConfigModel(sqlx.NewSqlConnFromSession(session))
}
