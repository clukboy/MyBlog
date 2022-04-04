// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	sq "github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	bizArticleLoveFieldNames          = builder.RawFieldNames(&BizArticleLove{}, true)
	bizArticleLoveRows                = strings.Join(bizArticleLoveFieldNames, ",")
	bizArticleLoveRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(bizArticleLoveFieldNames, "id", "create_time"))

	cachePublicBizArticleLoveIdPrefix = "cache:public:bizArticleLove:id:"
)

type (
	bizArticleLoveModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *BizArticleLove) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BizArticleLove, error)
		Update(ctx context.Context, session sqlx.Session, data *BizArticleLove) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultBizArticleLoveModel struct {
		sqlc.CachedConn
		table string
	}

	BizArticleLove struct {
		Id         int64          `db:"id"`
		ArticleId  int64          `db:"article_id"`  // 文章ID
		UserId     sql.NullInt64  `db:"user_id"`     // 已登录用户ID
		UserIp     sql.NullString `db:"user_ip"`     // 用户IP
		LoveTime   sql.NullTime   `db:"love_time"`   // 浏览时间
		CreateTime sql.NullTime   `db:"create_time"` // 添加时间
		UpdateTime sql.NullTime   `db:"update_time"` // 更新时间
	}
)

func newBizArticleLoveModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBizArticleLoveModel {
	return &defaultBizArticleLoveModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."biz_article_love"`,
	}
}

func (m *defaultBizArticleLoveModel) Insert(ctx context.Context, session sqlx.Session, data *BizArticleLove) (sql.Result, error) {
	query, args, err := sq.Insert(m.table).Columns(bizArticleLoveRows).Values(data.Id, data.ArticleId, data.UserId, data.UserIp, data.LoveTime, data.CreateTime, data.UpdateTime).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	publicBizArticleLoveIdKey := fmt.Sprintf("%s%v", cachePublicBizArticleLoveIdPrefix, data.Id)

	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicBizArticleLoveIdKey)
	return ret, err
}

func (m *defaultBizArticleLoveModel) FindOne(ctx context.Context, id int64) (*BizArticleLove, error) {
	publicBizArticleLoveIdKey := fmt.Sprintf("%s%v", cachePublicBizArticleLoveIdPrefix, id)
	var resp BizArticleLove
	err := m.QueryRowCtx(ctx, &resp, publicBizArticleLoveIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", bizArticleLoveRows, m.table)
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

func (m *defaultBizArticleLoveModel) Update(ctx context.Context, session sqlx.Session, data *BizArticleLove) error {
	publicBizArticleLoveIdKey := fmt.Sprintf("%s%v", cachePublicBizArticleLoveIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, bizArticleLoveRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.ArticleId, data.UserId, data.UserIp, data.LoveTime, time.Now())

		}
		return conn.ExecCtx(ctx, query, data.Id, data.ArticleId, data.UserId, data.UserIp, data.LoveTime, time.Now())
	}, publicBizArticleLoveIdKey)
	return err
}

func (m *defaultBizArticleLoveModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query, args, err := sq.Delete(m.table).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	publicBizArticleLoveIdKey := fmt.Sprintf("%s%v", cachePublicBizArticleLoveIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicBizArticleLoveIdKey)
	return err
}

func (m *defaultBizArticleLoveModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicBizArticleLoveIdPrefix, primary)
}

func (m *defaultBizArticleLoveModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", bizArticleLoveRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBizArticleLoveModel) tableName() string {
	return m.table
}
