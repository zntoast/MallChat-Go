package model

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserModel = (*customUserModel)(nil)

type (
	// UserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserModel.
	UserModel interface {
		userModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		MaxBuilder(field string) squirrel.SelectBuilder
		MinBuilder(field string) squirrel.SelectBuilder

		FindOneByQuery(ctx context.Context, session sqlx.Session, rowBuilder squirrel.SelectBuilder) (*User, error)
		FindSum(ctx context.Context, session sqlx.Session, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, session sqlx.Session, countBuilder squirrel.SelectBuilder) (int64, error)
		FindMaxInt64(ctx context.Context, session sqlx.Session, sumBuilder squirrel.SelectBuilder) (int64, error)
		FindMinInt64(ctx context.Context, session sqlx.Session, sumBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, session sqlx.Session, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*User, error)
		FindPageListByPage(ctx context.Context, session sqlx.Session, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*User, error)
		FindPageListByIdDESC(ctx context.Context, session sqlx.Session, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*User, error)
		FindPageListByIdASC(ctx context.Context, session sqlx.Session, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*User, error)
	}

	customUserModel struct {
		*defaultUserModel
	}
)

// NewUserModel returns a model for the database table.
func NewUserModel(conn sqlx.SqlConn, c cache.CacheConf) UserModel {
	return &customUserModel{
		defaultUserModel: newUserModel(conn, c),
	}
}

func (m *defaultUserModel) FindOneByQuery(ctx context.Context, session sqlx.Session, rowBuilder squirrel.SelectBuilder) (*User, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp User

	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, query, values...)
	} else {
		err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
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

func (m *defaultUserModel) FindSum(ctx context.Context, session sqlx.Session, sumBuilder squirrel.SelectBuilder) (float64, error) {

	query, values, err := sumBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64

	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, query, values...)
	} else {
		err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	}

	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultUserModel) FindCount(ctx context.Context, session sqlx.Session, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64

	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, query, values...)
	} else {
		err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	}

	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultUserModel) FindMaxInt64(ctx context.Context, session sqlx.Session, maxBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := maxBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64

	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, query, values...)
	} else {
		err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	}

	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultUserModel) FindMinInt64(ctx context.Context, session sqlx.Session, minBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := minBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64

	if session != nil {
		err = session.QueryRowCtx(ctx, &resp, query, values...)
	} else {
		err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	}

	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultUserModel) FindAll(ctx context.Context, session sqlx.Session, rowBuilder squirrel.SelectBuilder, orderBy string) ([]*User, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User

	if session != nil {
		err = session.QueryRowsCtx(ctx, &resp, query, values...)
	} else {
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	}

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindPageListByPage(ctx context.Context, session sqlx.Session, rowBuilder squirrel.SelectBuilder, page, pageSize int64, orderBy string) ([]*User, error) {

	if orderBy == "" {
		rowBuilder = rowBuilder.OrderBy("id DESC")
	} else {
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1 {
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User

	if session != nil {
		err = session.QueryRowsCtx(ctx, &resp, query, values...)
	} else {
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	}

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserModel) FindPageListByIdDESC(ctx context.Context, session sqlx.Session, rowBuilder squirrel.SelectBuilder, preMinId, pageSize int64) ([]*User, error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? ", preMinId)
	}

	query, values, err := rowBuilder.OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User

	if session != nil {
		err = session.QueryRowsCtx(ctx, &resp, query, values...)
	} else {
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	}

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// 按照id升序分页查询数据，不支持排序
func (m *defaultUserModel) FindPageListByIdASC(ctx context.Context, session sqlx.Session, rowBuilder squirrel.SelectBuilder, preMaxId, pageSize int64) ([]*User, error) {

	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? ", preMaxId)
	}

	query, values, err := rowBuilder.OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*User

	if session != nil {
		err = session.QueryRowsCtx(ctx, &resp, query, values...)
	} else {
		err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	}

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultUserModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultUserModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(userRows).From(m.table)
}

// export logic
func (m *defaultUserModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultUserModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

// export logic
func (m *defaultUserModel) MaxBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("MAX(" + field + ")").From(m.table)
}

// export logic
func (m *defaultUserModel) MinBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("MIN(" + field + ")").From(m.table)
}
