
func (m *default{{.upperStartCamelObject}}Model) Delete(ctx context.Context,session sqlx.Session, {{.lowerStartCamelPrimaryKey}} {{.dataType}}) error {
    query, args, err := sq.Delete(m.table).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
    if err != nil {
        return  err
    }
	{{if .withCache}}{{if .containsIndexCache}}data, err:=m.FindOne(ctx, {{.lowerStartCamelPrimaryKey}})
	if err!=nil{
		return err
	}
{{end}}	{{.keys}}
    _, err {{if .containsIndexCache}}={{else}}={{end}} m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
	if session != nil {
    return session.ExecCtx(ctx, query, args...)
    }
    return conn.ExecCtx(ctx, query, args...)
	}, {{.keyValues}}){{else}}
    if session != nil {
    _, err =  session.ExecCtx(ctx, query, args...)
    return err
    }
	_,err = m.conn.ExecCtx(ctx, query, {{.lowerStartCamelPrimaryKey}}){{end}}
	return err
}
