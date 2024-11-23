package {{.pkg}}
{{if .withCache}}
import (
	"context"
{{/*	"time"*/}}


	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/Masterminds/squirrel"
	{{/* "github.com/pkg/errors" */}}
)
{{else}}
import (
	"context"
{{/*	"time"*/}}

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/Masterminds/squirrel"
	{{/* "github.com/pkg/errors" */}}
)

{{end}}
var _ {{.upperStartCamelObject}}Model = (*custom{{.upperStartCamelObject}}Model)(nil)

type (
	// {{.upperStartCamelObject}}Model is an interface to be customized, add more methods here,
	// and implement the added methods in custom{{.upperStartCamelObject}}Model.
	{{.upperStartCamelObject}}Model interface {
		{{.lowerStartCamelObject}}Model
		Trans(ctx context.Context,fn func(context context.Context,session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		MaxBuilder(field string) squirrel.SelectBuilder
		MinBuilder(field string) squirrel.SelectBuilder
{{/*
		DeleteSoft(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}) error
*/}}
		FindOneByQuery(ctx context.Context,session sqlx.Session,rowBuilder squirrel.SelectBuilder) (*{{.upperStartCamelObject}},error)
		FindSum(ctx context.Context,session sqlx.Session,sumBuilder squirrel.SelectBuilder) (float64,error)
		FindCount(ctx context.Context,session sqlx.Session,countBuilder squirrel.SelectBuilder) (int64,error)
		FindMaxInt64(ctx context.Context,session sqlx.Session,sumBuilder squirrel.SelectBuilder) (int64,error)
		FindMinInt64(ctx context.Context,session sqlx.Session,sumBuilder squirrel.SelectBuilder) (int64,error)
		FindAll(ctx context.Context,session sqlx.Session,rowBuilder squirrel.SelectBuilder,orderBy string) ([]*{{.upperStartCamelObject}},error)
		FindPageListByPage(ctx context.Context,session sqlx.Session,rowBuilder squirrel.SelectBuilder,page ,pageSize int64,orderBy string) ([]*{{.upperStartCamelObject}},error)
		FindPageListByIdDESC(ctx context.Context,session sqlx.Session,rowBuilder squirrel.SelectBuilder ,preMinId ,pageSize int64) ([]*{{.upperStartCamelObject}},error)
		FindPageListByIdASC(ctx context.Context,session sqlx.Session,rowBuilder squirrel.SelectBuilder,preMaxId ,pageSize int64) ([]*{{.upperStartCamelObject}},error)

}

	custom{{.upperStartCamelObject}}Model struct {
		*default{{.upperStartCamelObject}}Model
	}
)

// New{{.upperStartCamelObject}}Model returns a model for the database table.
func New{{.upperStartCamelObject}}Model(conn sqlx.SqlConn{{if .withCache}}, c cache.CacheConf{{end}}) {{.upperStartCamelObject}}Model {
	return &custom{{.upperStartCamelObject}}Model{
		default{{.upperStartCamelObject}}Model: new{{.upperStartCamelObject}}Model(conn{{if .withCache}}, c{{end}}),
	}
}

{{/*
func (m *default{{.upperStartCamelObject}}Model) DeleteSoft(ctx context.Context,session sqlx.Session,data *{{.upperStartCamelObject}}) error {
	data.DelState = constant.DelStateYes
	data.DeleteTime = time.Now()
	if err:= m.UpdateWithVersion(ctx,session, data);err!= nil{
		return errors.Wrapf(xerr.NewSysErr("删除数据失败"),"{{.upperStartCamelObject}}Model delete err : %+v",err)
	}
	return nil
}
*/}}

func (m *default{{.upperStartCamelObject}}Model) FindOneByQuery(ctx context.Context,session sqlx.Session,rowBuilder squirrel.SelectBuilder) (*{{.upperStartCamelObject}},error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp {{.upperStartCamelObject}}
	{{if .withCache}}
		if session != nil{
			err = session.QueryRowCtx(ctx,&resp, query, values...)
		} else {
			err = m.QueryRowNoCacheCtx(ctx,&resp, query, values...)
		}
		
	{{else}}
		if session != nil{
			err = session.QueryRowCtx(ctx,&resp, query, values...)
		} else {
			err = m.conn.QueryRowCtx(ctx,&resp, query, values...)
		}
	{{end}}
	switch err {
	case nil:
		return &resp, nil
	case sqlx.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}


func (m *default{{.upperStartCamelObject}}Model) FindSum(ctx context.Context,session sqlx.Session,sumBuilder squirrel.SelectBuilder) (float64,error) {

	query, values, err := sumBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	{{if .withCache}}

		if session != nil{
			err = session.QueryRowCtx(ctx,&resp, query, values...)
		} else {
			err = m.QueryRowNoCacheCtx(ctx,&resp, query, values...)
		}
	
	{{else}}
		if session != nil{
			err = session.QueryRowCtx(ctx,&resp, query, values...)
		} else {
			err = m.conn.QueryRowCtx(ctx,&resp, query, values...)
		}
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) FindCount(ctx context.Context,session sqlx.Session,countBuilder squirrel.SelectBuilder) (int64,error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	{{if .withCache}}
		if session != nil{
			err = session.QueryRowCtx(ctx,&resp, query, values...)
		} else {
			err = m.QueryRowNoCacheCtx(ctx,&resp, query, values...)
		}
	
	{{else}}
		if session != nil{
			err = session.QueryRowCtx(ctx,&resp, query, values...)
		} else {
			err = m.conn.QueryRowCtx(ctx,&resp, query, values...)
		}
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}


func (m *default{{.upperStartCamelObject}}Model) FindMaxInt64(ctx context.Context,session sqlx.Session,maxBuilder squirrel.SelectBuilder) (int64,error) {

	query, values, err := maxBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	{{if .withCache}}
		if session != nil{
			err = session.QueryRowCtx(ctx,&resp, query, values...)
		} else {
			err = m.QueryRowNoCacheCtx(ctx,&resp, query, values...)
		}
	
	{{else}}
		if session != nil{
			err = session.QueryRowCtx(ctx,&resp, query, values...)
		} else {
			err = m.conn.QueryRowCtx(ctx,&resp, query, values...)
		}

	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) FindMinInt64(ctx context.Context,session sqlx.Session,minBuilder squirrel.SelectBuilder) (int64,error) {

	query, values, err := minBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	{{if .withCache}}
	
		if session != nil{
			err = session.QueryRowCtx(ctx,&resp, query, values...)
		} else {
			err = m.QueryRowNoCacheCtx(ctx,&resp, query, values...)
		}
		
	{{else}}
		if session != nil{
			err = session.QueryRowCtx(ctx,&resp, query, values...)
		} else {
			err = m.conn.QueryRowCtx(ctx,&resp, query, values...)
		}
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) FindAll(ctx context.Context,session sqlx.Session,rowBuilder squirrel.SelectBuilder,orderBy string) ([]*{{.upperStartCamelObject}},error) {

	if orderBy == ""{
		rowBuilder = rowBuilder.OrderBy("id DESC")
	}else{
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	{{if .withCache}}

		if session != nil{
			err = session.QueryRowsCtx(ctx,&resp, query, values...)
		} else {
			err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...)
		}
	{{else}}
		if session != nil{
			err = session.QueryRowsCtx(ctx,&resp, query, values...)
		} else {
			err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
		}
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) FindPageListByPage(ctx context.Context,session sqlx.Session,rowBuilder squirrel.SelectBuilder,page ,pageSize int64,orderBy string) ([]*{{.upperStartCamelObject}},error) {

	if orderBy == ""{
		rowBuilder = rowBuilder.OrderBy("id DESC")
	}else{
		rowBuilder = rowBuilder.OrderBy(orderBy)
	}

	if page < 1{
		page = 1
	}
	offset := (page - 1) * pageSize

	query, values, err := rowBuilder.Offset(uint64(offset)).Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	{{if .withCache}}
		if session != nil{
			err = session.QueryRowsCtx(ctx,&resp, query, values...)
		} else {
			err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...)
		}
	{{else}}
		if session != nil{
			err = session.QueryRowsCtx(ctx,&resp, query, values...)
		} else {
			err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
		}
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *default{{.upperStartCamelObject}}Model) FindPageListByIdDESC(ctx context.Context,session sqlx.Session,rowBuilder squirrel.SelectBuilder ,preMinId ,pageSize int64) ([]*{{.upperStartCamelObject}},error) {

	if preMinId > 0 {
		rowBuilder = rowBuilder.Where(" id < ? " , preMinId)
	}

	query, values, err := rowBuilder.OrderBy("id DESC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}
	

	{{if .withCache}}
		if session != nil{
			err = session.QueryRowsCtx(ctx,&resp, query, values...)
		} else {
			err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...)
		}
	{{else}}
		if session != nil{
			err = session.QueryRowsCtx(ctx,&resp, query, values...)
		} else {
			err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
		}
	{{end}}
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

//按照id升序分页查询数据，不支持排序
func (m *default{{.upperStartCamelObject}}Model) FindPageListByIdASC(ctx context.Context,session sqlx.Session,rowBuilder squirrel.SelectBuilder,preMaxId ,pageSize int64) ([]*{{.upperStartCamelObject}},error)  {

	if preMaxId > 0 {
		rowBuilder = rowBuilder.Where(" id > ? " , preMaxId)
	}

	query, values, err := rowBuilder.OrderBy("id ASC").Limit(uint64(pageSize)).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*{{.upperStartCamelObject}}

	{{if .withCache}}
		if session != nil{
			err = session.QueryRowsCtx(ctx,&resp, query, values...)
		} else {
			err = m.QueryRowsNoCacheCtx(ctx,&resp, query, values...)
		}
	{{else}}
		if session != nil{
			err = session.QueryRowsCtx(ctx,&resp, query, values...)
		} else {
			err = m.conn.QueryRowsCtx(ctx,&resp, query, values...)
		}
	{{end}}


	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *default{{.upperStartCamelObject}}Model) Trans(ctx context.Context,fn func(ctx context.Context,session sqlx.Session) error) error {
	{{if .withCache}}
	return m.TransactCtx(ctx,func(ctx context.Context,session sqlx.Session) error {
		return  fn(ctx,session)
	})
	{{else}}
	return m.conn.TransactCtx(ctx,func(ctx context.Context,session sqlx.Session) error {
		return  fn(ctx,session)
	})
	{{end}}
}

// export logic
func (m *default{{.upperStartCamelObject}}Model) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select({{.lowerStartCamelObject}}Rows).From(m.table)
}

// export logic
func (m *default{{.upperStartCamelObject}}Model) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT("+field+")").From(m.table)
}

// export logic
func (m *default{{.upperStartCamelObject}}Model) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM("+field+"),0)").From(m.table)
}

// export logic
func (m *default{{.upperStartCamelObject}}Model) MaxBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("MAX("+field+")").From(m.table)
}

// export logic
func (m *default{{.upperStartCamelObject}}Model) MinBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("MIN("+field+")").From(m.table)
}
