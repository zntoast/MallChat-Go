package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		withSession(session sqlx.Session) UserModel
		FindOneByQuery(ctx context.Context, session sqlx.Session, rowBuilder squirrel.SelectBuilder) (*User, error)
		RowBuilder() squirrel.SelectBuilder
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn),
	}
}

func (m *customUserModel) withSession(session sqlx.Session) UserModel {
	return NewUserModel(sqlx.NewSqlConnFromSession(session))
}

func (m *customUserModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select("*").From(m.table)
}

func (m *customUserModel) Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error {
	return m.conn.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *customUserModel) FindOneByQuery(ctx context.Context, session sqlx.Session, rowBuilder squirrel.SelectBuilder) (*User, error) {
	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp User
	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, query, values...)
	} else {
		err = m.conn.QueryRowCtx(ctx, &resp, query, values...)
	}

	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}
