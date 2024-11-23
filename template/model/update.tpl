
func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}})  (sql.Result,error) {
	{{if .withCache}}{{.keys}}

	newData := data
	_ = newData.Id

	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
		return session.ExecCtx(ctx,query, {{.expressionValues}})
	}
	return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
		return session.ExecCtx(ctx,query, {{.expressionValues}})
	}
	return m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
}




func (m *default{{.upperStartCamelObject}}Model) UpdateAndSql(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}, andSql string)  (sql.Result,error) {
	{{if .withCache}}{{.keys}}

	newData := data
	_ = newData.Id

	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} %s", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder, andSql)
	if session != nil{
		return session.ExecCtx(ctx,query, {{.expressionValues}})
	}
	return conn.ExecCtx(ctx, query, {{.expressionValues}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} %s", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder, andSql)
	if session != nil{
		return session.ExecCtx(ctx,query, {{.expressionValues}})
	}
	return m.conn.ExecCtx(ctx, query, {{.expressionValues}}){{end}}
}


func (m *default{{.upperStartCamelObject}}Model) IncInt64(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}, fieldName string, number int64)  (sql.Result,error) {
	{{if .withCache}}{{.keys}}


	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("update %s set %s = %s + ? where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, fieldName, fieldName)
	if session != nil{
		return session.ExecCtx(ctx,query, number, data.{{.upperStartCamelPrimaryKey}})
	}
	return conn.ExecCtx(ctx, query, number, data.{{.upperStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s = %s + ? where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table,  fieldName, fieldName)
	if session != nil{
		return session.ExecCtx(ctx,query, number, data.{{.upperStartCamelPrimaryKey}})
	}
	return m.conn.ExecCtx(ctx, query, number, data.{{.upperStartCamelPrimaryKey}}){{end}}
}


func (m *default{{.upperStartCamelObject}}Model) IncInt64AndSql(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}, fieldName string, number int64 ,andSql string)  (sql.Result,error) {
	{{if .withCache}}{{.keys}}

	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("update %s set %s = %s + ? where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} %s", m.table, fieldName, fieldName, andSql)
	if session != nil{
		return session.ExecCtx(ctx,query, number, data.{{.upperStartCamelPrimaryKey}})
	}
	return conn.ExecCtx(ctx, query, number, data.{{.upperStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s = %s + ? where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} %s", m.table, fieldName, fieldName, andSql)
	if session != nil{
		return session.ExecCtx(ctx,query, number, data.{{.upperStartCamelPrimaryKey}})
	}
	return m.conn.ExecCtx(ctx, query, number, data.{{.upperStartCamelPrimaryKey}}){{end}}
}

func (m *default{{.upperStartCamelObject}}Model) IncFloat64(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}, fieldName string, number float64)  (sql.Result,error) {
	{{if .withCache}}{{.keys}}


	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("update %s set %s = %s + ? where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, fieldName, fieldName)
	if session != nil{
		return session.ExecCtx(ctx,query, number, data.{{.upperStartCamelPrimaryKey}})
	}
	return conn.ExecCtx(ctx, query, number, data.{{.upperStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s = %s + ? where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table,  fieldName, fieldName)
	if session != nil{
		return session.ExecCtx(ctx,query, number, data.{{.upperStartCamelPrimaryKey}})
	}
	return m.conn.ExecCtx(ctx, query, number, data.{{.upperStartCamelPrimaryKey}}){{end}}
}


func (m *default{{.upperStartCamelObject}}Model) IncFloat64AndSql(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}, fieldName string, number float64 ,andSql string)  (sql.Result,error) {
	{{if .withCache}}{{.keys}}

	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("update %s set %s = %s + ? where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} %s", m.table, fieldName, fieldName, andSql)
	if session != nil{
		return session.ExecCtx(ctx,query, number, data.{{.upperStartCamelPrimaryKey}})
	}
	return conn.ExecCtx(ctx, query, number, data.{{.upperStartCamelPrimaryKey}})
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s = %s + ? where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} %s", m.table, fieldName, fieldName, andSql)
	if session != nil{
		return session.ExecCtx(ctx,query, number, data.{{.upperStartCamelPrimaryKey}})
	}
	return m.conn.ExecCtx(ctx, query, number, data.{{.upperStartCamelPrimaryKey}}){{end}}
}







{{/*
func (m *default{{.upperStartCamelObject}}Model) UpdateWithVersion(ctx context.Context,session sqlx.Session,data *{{.upperStartCamelObject}}) error {

	oldVersion := data.Version
	data.Version += 1

	var sqlResult sql.Result
	var err error

	{{if .withCache}}{{.keys}}
	sqlResult,err =  m.ExecCtx(ctx,func(ctx context.Context,conn sqlx.SqlConn) (result sql.Result, err error) {
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} and version = ? ", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
		return session.ExecCtx(ctx,query, {{.expressionValues}},oldVersion)
	}
	return conn.ExecCtx(ctx,query, {{.expressionValues}},oldVersion)
	}, {{.keyValues}}){{else}}query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}} and version = ? ", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
	if session != nil{
		sqlResult,err  =  session.ExecCtx(ctx,query, {{.expressionValues}},oldVersion)
	}else{
		sqlResult,err  =  m.conn.ExecCtx(ctx,query, {{.expressionValues}},oldVersion)
	}
	{{end}}
	if err != nil {
		return err
	}
	updateCount , err := sqlResult.RowsAffected()
	if err != nil{
		return err
	}
	if updateCount == 0 {
		return  xerr.NewErrCode(xerr.DbUpdateAffectedZeroError)
	}

	return nil
}
*/}}
