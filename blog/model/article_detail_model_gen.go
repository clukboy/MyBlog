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
	articleDetailFieldNames          = builder.RawFieldNames(&ArticleDetail{}, true)
	articleDetailRows                = strings.Join(articleDetailFieldNames, ",")
	articleDetailRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(articleDetailFieldNames, "id", "create_time"))

	cachePublicArticleDetailIdPrefix = "cache:public:articleDetail:id:"
)

type (
	articleDetailModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *ArticleDetail) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ArticleDetail, error)
		Update(ctx context.Context, session sqlx.Session, data *ArticleDetail) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultArticleDetailModel struct {
		sqlc.CachedConn
		table string
	}

	ArticleDetail struct {
		Id          int64     `db:"id"`           // 文章详情id
		ContentMd   string    `db:"content_md"`   // 文章markdown内容
		ContentHtml string    `db:"content_html"` // 文章html内容
		ArticleId   int64     `db:"article_id"`   // 文章id
		CreateTime  time.Time `db:"create_time"`  // 创建时间
		UpdateTime  time.Time `db:"update_time"`  // 修改时间
	}
)

func newArticleDetailModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultArticleDetailModel {
	return &defaultArticleDetailModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."article_detail"`,
	}
}

func (m *defaultArticleDetailModel) Insert(ctx context.Context, session sqlx.Session, data *ArticleDetail) (sql.Result, error) {
	query, args, err := sq.Insert(m.table).Columns(articleDetailRows).Values(data.Id, data.ContentMd, data.ContentHtml, data.ArticleId, data.CreateTime, data.UpdateTime).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	publicArticleDetailIdKey := fmt.Sprintf("%s%v", cachePublicArticleDetailIdPrefix, data.Id)

	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicArticleDetailIdKey)
	return ret, err
}

func (m *defaultArticleDetailModel) FindOne(ctx context.Context, id int64) (*ArticleDetail, error) {
	publicArticleDetailIdKey := fmt.Sprintf("%s%v", cachePublicArticleDetailIdPrefix, id)
	var resp ArticleDetail
	err := m.QueryRowCtx(ctx, &resp, publicArticleDetailIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", articleDetailRows, m.table)
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

func (m *defaultArticleDetailModel) Update(ctx context.Context, session sqlx.Session, data *ArticleDetail) error {
	publicArticleDetailIdKey := fmt.Sprintf("%s%v", cachePublicArticleDetailIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, articleDetailRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.ContentMd, data.ContentHtml, data.ArticleId, time.Now())

		}
		return conn.ExecCtx(ctx, query, data.Id, data.ContentMd, data.ContentHtml, data.ArticleId, time.Now())
	}, publicArticleDetailIdKey)
	return err
}

func (m *defaultArticleDetailModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query, args, err := sq.Delete(m.table).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	publicArticleDetailIdKey := fmt.Sprintf("%s%v", cachePublicArticleDetailIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicArticleDetailIdKey)
	return err
}

func (m *defaultArticleDetailModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicArticleDetailIdPrefix, primary)
}

func (m *defaultArticleDetailModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", articleDetailRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultArticleDetailModel) tableName() string {
	return m.table
}
