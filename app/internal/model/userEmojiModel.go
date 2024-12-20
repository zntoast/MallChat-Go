package model

import "github.com/zeromicro/go-zero/core/stores/sqlx"

var _ UserEmojiModel = (*customUserEmojiModel)(nil)

type (
	// UserEmojiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserEmojiModel.
	UserEmojiModel interface {
		userEmojiModel
		withSession(session sqlx.Session) UserEmojiModel
	}

	customUserEmojiModel struct {
		*defaultUserEmojiModel
	}
)

// NewUserEmojiModel returns a model for the database table.
func NewUserEmojiModel(conn sqlx.SqlConn) UserEmojiModel {
	return &customUserEmojiModel{
		defaultUserEmojiModel: newUserEmojiModel(conn),
	}
}

func (m *customUserEmojiModel) withSession(session sqlx.Session) UserEmojiModel {
	return NewUserEmojiModel(sqlx.NewSqlConnFromSession(session))
}
