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
	articleFieldNames          = builder.RawFieldNames(&Article{}, true)
	articleRows                = strings.Join(articleFieldNames, ",")
	articleRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(articleFieldNames, "id", "create_time"))

	cachePublicArticleIdPrefix = "cache:public:article:id:"
)

type (
	articleModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Article) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Article, error)
		Update(ctx context.Context, session sqlx.Session, data *Article) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultArticleModel struct {
		sqlc.CachedConn
		table string
	}

	Article struct {
		Id             int64     `db:"id"`              // 文章id
		PushDate       time.Time `db:"push_date"`       // 发布日期
		Uid            int64     `db:"uid"`             // 发表用户id
		Title          string    `db:"title"`           // 博文标题
		LikeCount      int64     `db:"like_count"`      // 点赞数
		CommentCount   int64     `db:"comment_count"`   // 评论数
		ReadCount      int64     `db:"read_count"`      // 阅读数
		TopFlag        bool      `db:"top_flag"`        // 是否顶置
		ArticleSummary string    `db:"article_summary"` // 文章摘要
		CreateTime     time.Time `db:"create_time"`     // 创建时间
		UpdateTime     time.Time `db:"update_time"`     // 修改时间
	}
)

func newArticleModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultArticleModel {
	return &defaultArticleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."article"`,
	}
}

func (m *defaultArticleModel) Insert(ctx context.Context, session sqlx.Session, data *Article) (sql.Result, error) {
	query, args, err := sq.Insert(m.table).Columns(articleRows).Values(data.Id, data.PushDate, data.Uid, data.Title, data.LikeCount, data.CommentCount, data.ReadCount, data.TopFlag, data.ArticleSummary, data.CreateTime, data.UpdateTime).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	publicArticleIdKey := fmt.Sprintf("%s%v", cachePublicArticleIdPrefix, data.Id)

	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicArticleIdKey)
	return ret, err
}

func (m *defaultArticleModel) FindOne(ctx context.Context, id int64) (*Article, error) {
	publicArticleIdKey := fmt.Sprintf("%s%v", cachePublicArticleIdPrefix, id)
	var resp Article
	err := m.QueryRowCtx(ctx, &resp, publicArticleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", articleRows, m.table)
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

func (m *defaultArticleModel) Update(ctx context.Context, session sqlx.Session, data *Article) error {
	publicArticleIdKey := fmt.Sprintf("%s%v", cachePublicArticleIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, articleRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.PushDate, data.Uid, data.Title, data.LikeCount, data.CommentCount, data.ReadCount, data.TopFlag, data.ArticleSummary, time.Now())

		}
		return conn.ExecCtx(ctx, query, data.Id, data.PushDate, data.Uid, data.Title, data.LikeCount, data.CommentCount, data.ReadCount, data.TopFlag, data.ArticleSummary, time.Now())
	}, publicArticleIdKey)
	return err
}

func (m *defaultArticleModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query, args, err := sq.Delete(m.table).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	publicArticleIdKey := fmt.Sprintf("%s%v", cachePublicArticleIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicArticleIdKey)
	return err
}

func (m *defaultArticleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicArticleIdPrefix, primary)
}

func (m *defaultArticleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", articleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultArticleModel) tableName() string {
	return m.table
}
