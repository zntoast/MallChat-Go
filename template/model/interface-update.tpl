Update(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result, error)
UpdateAndSql(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}, andSql string) (sql.Result, error)
IncInt64(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}, fieldName string, number int64)  (sql.Result,error)
IncInt64AndSql(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}, fieldName string, number int64 ,andSql string)  (sql.Result,error)
IncFloat64(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}, fieldName string, number float64)  (sql.Result,error)
IncFloat64AndSql(ctx context.Context,session sqlx.Session, data *{{.upperStartCamelObject}}, fieldName string, number float64 ,andSql string)  (sql.Result,error)
{{/*UpdateWithVersion(ctx context.Context,session sqlx.Session,data *{{.upperStartCamelObject}}) error*/}}