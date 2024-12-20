package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserApplyModel = (*customUserApplyModel)(nil)

type (
	// UserApplyModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserApplyModel.
	UserApplyModel interface {
		userApplyModel
		withSession(session sqlx.Session) UserApplyModel
	}

	customUserApplyModel struct {
		*defaultUserApplyModel
	}
)

// NewUserApplyModel returns a model for the database table.
func NewUserApplyModel(conn sqlx.SqlConn) UserApplyModel {
	return &customUserApplyModel{
		defaultUserApplyModel: newUserApplyModel(conn),
	}
}

func (m *customUserApplyModel) withSession(session sqlx.Session) UserApplyModel {
	return NewUserApplyModel(sqlx.NewSqlConnFromSession(session))
}
