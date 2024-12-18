// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.2

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	groupFieldNames          = builder.RawFieldNames(&Group{})
	groupRows                = strings.Join(groupFieldNames, ",")
	groupRowsExpectAutoSet   = strings.Join(stringx.Remove(groupFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	groupRowsWithPlaceHolder = strings.Join(stringx.Remove(groupFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheMallchatgoGroupIdPrefix = "cache:mallchatgo:group:id:"
)

type (
	groupModel interface {
		Insert(ctx context.Context, data *Group) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Group, error)
		Update(ctx context.Context, data *Group) error
		Delete(ctx context.Context, id int64) error
	}

	defaultGroupModel struct {
		sqlc.CachedConn
		table string
	}

	Group struct {
		Id         int64  `db:"id"`
		Name       string `db:"name"`        // 群名称
		Avatar     string `db:"avatar"`      // 群头像
		CreatorId  int64  `db:"creator_id"`  // 创建者ID
		CreateTime int64  `db:"create_time"` // 创建时间
		UpdateTime int64  `db:"update_time"` // 更新时间
	}
)

func newGroupModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultGroupModel {
	return &defaultGroupModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`group`",
	}
}

func (m *defaultGroupModel) Delete(ctx context.Context, id int64) error {
	mallchatgoGroupIdKey := fmt.Sprintf("%s%v", cacheMallchatgoGroupIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, mallchatgoGroupIdKey)
	return err
}

func (m *defaultGroupModel) FindOne(ctx context.Context, id int64) (*Group, error) {
	mallchatgoGroupIdKey := fmt.Sprintf("%s%v", cacheMallchatgoGroupIdPrefix, id)
	var resp Group
	err := m.QueryRowCtx(ctx, &resp, mallchatgoGroupIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGroupModel) Insert(ctx context.Context, data *Group) (sql.Result, error) {
	mallchatgoGroupIdKey := fmt.Sprintf("%s%v", cacheMallchatgoGroupIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?)", m.table, groupRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Avatar, data.CreatorId)
	}, mallchatgoGroupIdKey)
	return ret, err
}

func (m *defaultGroupModel) Update(ctx context.Context, data *Group) error {
	mallchatgoGroupIdKey := fmt.Sprintf("%s%v", cacheMallchatgoGroupIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, groupRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Avatar, data.CreatorId, data.Id)
	}, mallchatgoGroupIdKey)
	return err
}

func (m *defaultGroupModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheMallchatgoGroupIdPrefix, primary)
}

func (m *defaultGroupModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultGroupModel) tableName() string {
	return m.table
}