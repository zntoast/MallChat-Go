package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserBackpackModel = (*customUserBackpackModel)(nil)

type (
	// UserBackpackModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserBackpackModel.
	UserBackpackModel interface {
		userBackpackModel
		withSession(session sqlx.Session) UserBackpackModel
	}

	customUserBackpackModel struct {
		*defaultUserBackpackModel
	}
)

// NewUserBackpackModel returns a model for the database table.
func NewUserBackpackModel(conn sqlx.SqlConn) UserBackpackModel {
	return &customUserBackpackModel{
		defaultUserBackpackModel: newUserBackpackModel(conn),
	}
}

func (m *customUserBackpackModel) withSession(session sqlx.Session) UserBackpackModel {
	return NewUserBackpackModel(sqlx.NewSqlConnFromSession(session))
}
