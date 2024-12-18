package model

import (
	"context"
	"database/sql"

	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		ExistsByMobile(ctx context.Context, mobile string) (bool, error)
		UpdatePassword(mobile string, newPassword string) error
		FindByMobile(mobile string) (*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c, opts...),
	}
}

func (m *customUserModel) ExistsByMobile(ctx context.Context, mobile string) (bool, error) {
	var count int64
	query := `select count(*) from user where mobile = ?`
	err := m.CachedConn.QueryRowNoCacheCtx(ctx, &count, query, mobile)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (m *customUserModel) UpdatePassword(mobile string, newPassword string) error {
	query := `UPDATE user SET password = ? WHERE mobile = ?`
	_, err := m.ExecCtx(context.Background(), func(ctx context.Context, conn sqlx.SqlConn) (sql.Result, error) {
		return conn.ExecCtx(ctx, query, newPassword, mobile)
	})
	return err
}

func (m *customUserModel) FindByMobile(mobile string) (*User, error) {
	var user User
	query := `SELECT * FROM user WHERE mobile = ?`
	err := m.CachedConn.QueryRowNoCacheCtx(context.Background(), &user, query, mobile)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
