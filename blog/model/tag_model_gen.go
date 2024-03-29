// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	tagFieldNames          = builder.RawFieldNames(&Tag{}, true)
	tagRows                = strings.Join(tagFieldNames, ",")
	tagRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(tagFieldNames, "id", "create_time"))

	cachePublicTagIdPrefix = "cache:public:tag:id:"
)

type (
	tagModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Tag) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Tag, error)
		Update(ctx context.Context, session sqlx.Session, data *Tag) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultTagModel struct {
		sqlc.CachedConn
		table string
	}

	Tag struct {
		Id          int64     `db:"id"`          // 标签id
		Name        string    `db:"name"`        // 标签名称
		AliasName   string    `db:"alias_name"`  // 标签别名
		Description string    `db:"description"` // 标签描述
		CreateTime  time.Time `db:"create_time"` // 创建时间
		UpdateTime  time.Time `db:"update_time"` // 修改时间
	}
)

func newTagModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultTagModel {
	return &defaultTagModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."tag"`,
	}
}

func (m *defaultTagModel) Insert(ctx context.Context, session sqlx.Session, data *Tag) (sql.Result, error) {
	query, args, err := sq.Insert(m.table).Columns(tagRows).Values(data.Id, data.Name, data.AliasName, data.Description, data.CreateTime, data.UpdateTime).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	publicTagIdKey := fmt.Sprintf("%s%v", cachePublicTagIdPrefix, data.Id)

	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicTagIdKey)
	return ret, err
}

func (m *defaultTagModel) FindOne(ctx context.Context, id int64) (*Tag, error) {
	publicTagIdKey := fmt.Sprintf("%s%v", cachePublicTagIdPrefix, id)
	var resp Tag
	err := m.QueryRowCtx(ctx, &resp, publicTagIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", tagRows, m.table)
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

func (m *defaultTagModel) Update(ctx context.Context, session sqlx.Session, data *Tag) error {
	publicTagIdKey := fmt.Sprintf("%s%v", cachePublicTagIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, tagRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.Name, data.AliasName, data.Description, time.Now())

		}
		return conn.ExecCtx(ctx, query, data.Id, data.Name, data.AliasName, data.Description, time.Now())
	}, publicTagIdKey)
	return err
}

func (m *defaultTagModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query, args, err := sq.Delete(m.table).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	publicTagIdKey := fmt.Sprintf("%s%v", cachePublicTagIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicTagIdKey)
	return err
}

func (m *defaultTagModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicTagIdPrefix, primary)
}

func (m *defaultTagModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", tagRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultTagModel) tableName() string {
	return m.table
}
