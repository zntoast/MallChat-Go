// Code generated by goctl. DO NOT EDIT!

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
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`"), "=?,") + "=?"

	cacheMallchatgoUserIdPrefix = "cache:mallchatgo:user:id:"
)

type (
	userModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		FindOne(ctx context.Context, session sqlx.Session, id int64) (*User, error)
		Update(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		UpdateAndSql(ctx context.Context, session sqlx.Session, data *User, andSql string) (sql.Result, error)
		IncInt64(ctx context.Context, session sqlx.Session, data *User, fieldName string, number int64) (sql.Result, error)
		IncInt64AndSql(ctx context.Context, session sqlx.Session, data *User, fieldName string, number int64, andSql string) (sql.Result, error)
		IncFloat64(ctx context.Context, session sqlx.Session, data *User, fieldName string, number float64) (sql.Result, error)
		IncFloat64AndSql(ctx context.Context, session sqlx.Session, data *User, fieldName string, number float64, andSql string) (sql.Result, error)

		Delete(ctx context.Context, session sqlx.Session, id int64) error
		DeleteAllTruncate(ctx context.Context, session sqlx.Session) error
		DeleteBatchById(ctx context.Context, session sqlx.Session, ids []int64) error
		DeleteBatchByIdAndSql(ctx context.Context, session sqlx.Session, ids []int64, andSql string) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id           int64          `db:"id"`            // 用户id
		Name         string         `db:"name"`          // 用户昵称
		Avatar       sql.NullString `db:"avatar"`        // 用户头像
		Sex          sql.NullInt64  `db:"sex"`           // 性别 1为男性，2为女性
		OpenId       sql.NullString `db:"open_id"`       // 微信openid用户标识
		ActiveStatus sql.NullInt64  `db:"active_status"` // 上下线状态 1在线 2离线
		LastOptTime  sql.NullTime   `db:"last_opt_time"` // 最后上下线时间
		IpInfo       sql.NullString `db:"ip_info"`       // IP信息
		ItemId       sql.NullInt64  `db:"item_id"`       // 佩戴的徽章id
		Status       sql.NullInt64  `db:"status"`        // 用户状态 0正常 1拉黑
		CreateTime   sql.NullTime   `db:"create_time"`   // 创建时间
		UpdateTime   sql.NullTime   `db:"update_time"`   // 修改时间
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	//data.DeleteTime = time.Unix(0,0)
	mallchatgoUserIdKey := fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, userRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Name, data.Avatar, data.Sex, data.OpenId, data.ActiveStatus, data.LastOptTime, data.IpInfo, data.ItemId, data.Status)
		}
		return conn.ExecCtx(ctx, query, data.Name, data.Avatar, data.Sex, data.OpenId, data.ActiveStatus, data.LastOptTime, data.IpInfo, data.ItemId, data.Status)
	}, mallchatgoUserIdKey)
}

func (m *defaultUserModel) FindOne(ctx context.Context, session sqlx.Session, id int64) (*User, error) {
	mallchatgoUserIdKey := fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, id)
	var resp User
	var err error

	if session != nil {
		query := fmt.Sprintf("select %s from %s where `id` = ?  limit 1", userRows, m.table)
		err = session.QueryRowCtx(ctx, &resp, query, id)
	} else {
		err = m.QueryRowCtx(ctx, &resp, mallchatgoUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
			query := fmt.Sprintf("select %s from %s where `id` = ?  limit 1", userRows, m.table)
			return conn.QueryRowCtx(ctx, v, query, id)
		})
	}

	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Update(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	mallchatgoUserIdKey := fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, data.Id)

	newData := data
	_ = newData.Id

	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Name, data.Avatar, data.Sex, data.OpenId, data.ActiveStatus, data.LastOptTime, data.IpInfo, data.ItemId, data.Status, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.Name, data.Avatar, data.Sex, data.OpenId, data.ActiveStatus, data.LastOptTime, data.IpInfo, data.ItemId, data.Status, data.Id)
	}, mallchatgoUserIdKey)
}

func (m *defaultUserModel) UpdateAndSql(ctx context.Context, session sqlx.Session, data *User, andSql string) (sql.Result, error) {
	mallchatgoUserIdKey := fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, data.Id)

	newData := data
	_ = newData.Id

	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ? %s", m.table, userRowsWithPlaceHolder, andSql)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Name, data.Avatar, data.Sex, data.OpenId, data.ActiveStatus, data.LastOptTime, data.IpInfo, data.ItemId, data.Status, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.Name, data.Avatar, data.Sex, data.OpenId, data.ActiveStatus, data.LastOptTime, data.IpInfo, data.ItemId, data.Status, data.Id)
	}, mallchatgoUserIdKey)
}

func (m *defaultUserModel) IncInt64(ctx context.Context, session sqlx.Session, data *User, fieldName string, number int64) (sql.Result, error) {
	mallchatgoUserIdKey := fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, data.Id)

	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s = %s + ? where `id` = ?", m.table, fieldName, fieldName)
		if session != nil {
			return session.ExecCtx(ctx, query, number, data.Id)
		}
		return conn.ExecCtx(ctx, query, number, data.Id)
	}, mallchatgoUserIdKey)
}

func (m *defaultUserModel) IncInt64AndSql(ctx context.Context, session sqlx.Session, data *User, fieldName string, number int64, andSql string) (sql.Result, error) {
	mallchatgoUserIdKey := fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, data.Id)

	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s = %s + ? where `id` = ? %s", m.table, fieldName, fieldName, andSql)
		if session != nil {
			return session.ExecCtx(ctx, query, number, data.Id)
		}
		return conn.ExecCtx(ctx, query, number, data.Id)
	}, mallchatgoUserIdKey)
}

func (m *defaultUserModel) IncFloat64(ctx context.Context, session sqlx.Session, data *User, fieldName string, number float64) (sql.Result, error) {
	mallchatgoUserIdKey := fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, data.Id)

	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s = %s + ? where `id` = ?", m.table, fieldName, fieldName)
		if session != nil {
			return session.ExecCtx(ctx, query, number, data.Id)
		}
		return conn.ExecCtx(ctx, query, number, data.Id)
	}, mallchatgoUserIdKey)
}

func (m *defaultUserModel) IncFloat64AndSql(ctx context.Context, session sqlx.Session, data *User, fieldName string, number float64, andSql string) (sql.Result, error) {
	mallchatgoUserIdKey := fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, data.Id)

	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s = %s + ? where `id` = ? %s", m.table, fieldName, fieldName, andSql)
		if session != nil {
			return session.ExecCtx(ctx, query, number, data.Id)
		}
		return conn.ExecCtx(ctx, query, number, data.Id)
	}, mallchatgoUserIdKey)
}

func (m *defaultUserModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	mallchatgoUserIdKey := fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, mallchatgoUserIdKey)
	return err
}

func (m *defaultUserModel) DeleteAllTruncate(ctx context.Context, session sqlx.Session) error {

	data := &User{}
	_ = data.Id

	var id = "*"
	mallchatgoUserIdKey := fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, id)

	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("truncate %s", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query)
		}
		return conn.ExecCtx(ctx, query)
	}, mallchatgoUserIdKey)
	return err
}
func (m *defaultUserModel) DeleteBatchById(ctx context.Context, session sqlx.Session, ids []int64) error {
	var res []string
	for _, v := range ids {
		res = append(res, fmt.Sprintf("'%v'", v))
	}
	delArr := strings.Join(res, ",")

	data := &User{}
	_ = data.Id

	var id = "*"
	mallchatgoUserIdKey := fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("DELETE FROM %s WHERE id in (%s)", m.table, delArr)
		if session != nil {
			return session.ExecCtx(ctx, query)
		}
		return conn.ExecCtx(ctx, query)
	}, mallchatgoUserIdKey)
	return err
}

func (m *defaultUserModel) DeleteBatchByIdAndSql(ctx context.Context, session sqlx.Session, ids []int64, andSql string) error {
	var res []string
	for _, v := range ids {
		res = append(res, fmt.Sprintf("'%v'", v))
	}
	delArr := strings.Join(res, ",")

	data := &User{}
	_ = data.Id

	var id = "*"
	mallchatgoUserIdKey := fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("DELETE FROM %s WHERE id in (%s) %s", m.table, delArr, andSql)
		if session != nil {
			return session.ExecCtx(ctx, query)
		}
		return conn.ExecCtx(ctx, query)
	}, mallchatgoUserIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheMallchatgoUserIdPrefix, primary)
}
func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
