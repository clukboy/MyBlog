
func (m *default{{.upperStartCamelObject}}Model) Update(ctx context.Context, session sqlx.Session, data *{{.upperStartCamelObject}}) error {
	{{if .withCache}}{{.keys}}
    _, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
		if session != nil {
            return session.ExecCtx(ctx, query, data.Id, data.Name, time.Now())

        }
		return conn.ExecCtx(ctx, query, {{.expressionValues}},time.Now())
	}, {{.keyValues}}){{else}}
	query := fmt.Sprintf("update %s set %s where {{.originalPrimaryKey}} = {{if .postgreSql}}$1{{else}}?{{end}}", m.table, {{.lowerStartCamelObject}}RowsWithPlaceHolder)
    	if session != nil {
    		_, err := session.ExecCtx(ctx, query, data.Id, data.Name, time.Now())
    		return err
    	}
    _,err:=m.conn.ExecCtx(ctx, query, {{.expressionValues}},time.Now()){{end}}
	return err
}
