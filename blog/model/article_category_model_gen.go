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
	articleCategoryFieldNames          = builder.RawFieldNames(&ArticleCategory{}, true)
	articleCategoryRows                = strings.Join(articleCategoryFieldNames, ",")
	articleCategoryRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(articleCategoryFieldNames, "id", "create_time"))

	cachePublicArticleCategoryIdPrefix = "cache:public:articleCategory:id:"
)

type (
	articleCategoryModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *ArticleCategory) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*ArticleCategory, error)
		Update(ctx context.Context, session sqlx.Session, data *ArticleCategory) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultArticleCategoryModel struct {
		sqlc.CachedConn
		table string
	}

	ArticleCategory struct {
		Id         int64     `db:"id"`          // 文章分类id
		ArticleId  int64     `db:"article_id"`  // 文章id
		CategoryId int64     `db:"category_id"` // 分类id
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 修改时间
	}
)

func newArticleCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultArticleCategoryModel {
	return &defaultArticleCategoryModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."article_category"`,
	}
}

func (m *defaultArticleCategoryModel) Insert(ctx context.Context, session sqlx.Session, data *ArticleCategory) (sql.Result, error) {
	query, args, err := sq.Insert(m.table).Columns(articleCategoryRows).Values(data.Id, data.ArticleId, data.CategoryId, data.CreateTime, data.UpdateTime).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	publicArticleCategoryIdKey := fmt.Sprintf("%s%v", cachePublicArticleCategoryIdPrefix, data.Id)

	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicArticleCategoryIdKey)
	return ret, err
}

func (m *defaultArticleCategoryModel) FindOne(ctx context.Context, id int64) (*ArticleCategory, error) {
	publicArticleCategoryIdKey := fmt.Sprintf("%s%v", cachePublicArticleCategoryIdPrefix, id)
	var resp ArticleCategory
	err := m.QueryRowCtx(ctx, &resp, publicArticleCategoryIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", articleCategoryRows, m.table)
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

func (m *defaultArticleCategoryModel) Update(ctx context.Context, session sqlx.Session, data *ArticleCategory) error {
	publicArticleCategoryIdKey := fmt.Sprintf("%s%v", cachePublicArticleCategoryIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, articleCategoryRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.ArticleId, data.CategoryId, time.Now())

		}
		return conn.ExecCtx(ctx, query, data.Id, data.ArticleId, data.CategoryId, time.Now())
	}, publicArticleCategoryIdKey)
	return err
}

func (m *defaultArticleCategoryModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query, args, err := sq.Delete(m.table).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	publicArticleCategoryIdKey := fmt.Sprintf("%s%v", cachePublicArticleCategoryIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicArticleCategoryIdKey)
	return err
}

func (m *defaultArticleCategoryModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicArticleCategoryIdPrefix, primary)
}

func (m *defaultArticleCategoryModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", articleCategoryRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultArticleCategoryModel) tableName() string {
	return m.table
}
