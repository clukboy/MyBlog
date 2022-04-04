
func (m *default{{.upperStartCamelObject}}Model) Insert(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) (sql.Result,error) {
	query, args, err := sq.Insert(m.table).Columns({{.lowerStartCamelObject}}Rows).Values({{.expressionValues}},data.CreateTime,data.UpdateTime).
                               PlaceholderFormat(sq.Dollar).ToSql()
    if err != nil {
        return nil, err
    }
	{{if .withCache}}{{.keys}}

    ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
    if session != nil {
        return session.ExecCtx(ctx, query, args...)
    }
        return conn.ExecCtx(ctx, query, args...)
	}, {{.keyValues}}){{else}}
    if session != nil {
        return session.ExecCtx(ctx, query, args...)
    }
    ret,err := m.conn.ExecCtx(ctx, query, args...)
	{{end}}
	return ret,err
}
