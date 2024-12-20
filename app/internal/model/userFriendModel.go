package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserFriendModel = (*customUserFriendModel)(nil)

type (
	// UserFriendModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserFriendModel.
	UserFriendModel interface {
		userFriendModel
		withSession(session sqlx.Session) UserFriendModel
	}

	customUserFriendModel struct {
		*defaultUserFriendModel
	}
)

// NewUserFriendModel returns a model for the database table.
func NewUserFriendModel(conn sqlx.SqlConn) UserFriendModel {
	return &customUserFriendModel{
		defaultUserFriendModel: newUserFriendModel(conn),
	}
}

func (m *customUserFriendModel) withSession(session sqlx.Session) UserFriendModel {
	return NewUserFriendModel(sqlx.NewSqlConnFromSession(session))
}
