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
	bizArticleFieldNames          = builder.RawFieldNames(&BizArticle{}, true)
	bizArticleRows                = strings.Join(bizArticleFieldNames, ",")
	bizArticleRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(bizArticleFieldNames, "id", "create_time"))

	cachePublicBizArticleIdPrefix = "cache:public:bizArticle:id:"
)

type (
	bizArticleModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *BizArticle) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*BizArticle, error)
		Update(ctx context.Context, session sqlx.Session, data *BizArticle) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultBizArticleModel struct {
		sqlc.CachedConn
		table string
	}

	BizArticle struct {
		Id           int64     `db:"id"`
		Title        string    `db:"title"`       // 文章标题
		UserId       int64     `db:"user_id"`     // 用户ID
		CoverImage   string    `db:"cover_image"` // 文章封面图片
		EditorType   string    `db:"editor_type"` // 当前文章适用的编辑器类型
		QrcodePath   string    `db:"qrcode_path"` // 文章专属二维码地址
		IsMarkdown   int64     `db:"is_markdown"`
		Content      string    `db:"content"`       // 文章内容
		ContentMd    string    `db:"content_md"`    // markdown版的文章内容
		Top          int64     `db:"top"`           // 是否置顶
		TypeId       int64     `db:"type_id"`       // 类型
		Status       int64     `db:"status"`        // 状态
		Recommended  int64     `db:"recommended"`   // 是否推荐
		Original     int64     `db:"original"`      // 是否原创
		Description  string    `db:"description"`   // 文章简介，最多200字
		Keywords     string    `db:"keywords"`      // 文章关键字，优化搜索
		Comment      int64     `db:"comment"`       // 是否开启评论
		Password     string    `db:"password"`      // 文章私密访问时的密钥
		RequiredAuth int64     `db:"required_auth"` // 该文章是否登录后才可访问
		CreateTime   time.Time `db:"create_time"`   // 添加时间
		UpdateTime   time.Time `db:"update_time"`   // 更新时间
	}
)

func newBizArticleModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultBizArticleModel {
	return &defaultBizArticleModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."biz_article"`,
	}
}

func (m *defaultBizArticleModel) Insert(ctx context.Context, session sqlx.Session, data *BizArticle) (sql.Result, error) {
	query, args, err := sq.Insert(m.table).Columns(bizArticleRows).Values(data.Id, data.Title, data.UserId, data.CoverImage, data.EditorType, data.QrcodePath, data.IsMarkdown, data.Content, data.ContentMd, data.Top, data.TypeId, data.Status, data.Recommended, data.Original, data.Description, data.Keywords, data.Comment, data.Password, data.RequiredAuth, data.CreateTime, data.UpdateTime).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	publicBizArticleIdKey := fmt.Sprintf("%s%v", cachePublicBizArticleIdPrefix, data.Id)

	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicBizArticleIdKey)
	return ret, err
}

func (m *defaultBizArticleModel) FindOne(ctx context.Context, id int64) (*BizArticle, error) {
	publicBizArticleIdKey := fmt.Sprintf("%s%v", cachePublicBizArticleIdPrefix, id)
	var resp BizArticle
	err := m.QueryRowCtx(ctx, &resp, publicBizArticleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", bizArticleRows, m.table)
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

func (m *defaultBizArticleModel) Update(ctx context.Context, session sqlx.Session, data *BizArticle) error {
	publicBizArticleIdKey := fmt.Sprintf("%s%v", cachePublicBizArticleIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, bizArticleRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.Title, data.UserId, data.CoverImage, data.EditorType, data.QrcodePath, data.IsMarkdown, data.Content, data.ContentMd, data.Top, data.TypeId, data.Status, data.Recommended, data.Original, data.Description, data.Keywords, data.Comment, data.Password, data.RequiredAuth, time.Now())

		}
		return conn.ExecCtx(ctx, query, data.Id, data.Title, data.UserId, data.CoverImage, data.EditorType, data.QrcodePath, data.IsMarkdown, data.Content, data.ContentMd, data.Top, data.TypeId, data.Status, data.Recommended, data.Original, data.Description, data.Keywords, data.Comment, data.Password, data.RequiredAuth, time.Now())
	}, publicBizArticleIdKey)
	return err
}

func (m *defaultBizArticleModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query, args, err := sq.Delete(m.table).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	publicBizArticleIdKey := fmt.Sprintf("%s%v", cachePublicBizArticleIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicBizArticleIdKey)
	return err
}

func (m *defaultBizArticleModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicBizArticleIdPrefix, primary)
}

func (m *defaultBizArticleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", bizArticleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultBizArticleModel) tableName() string {
	return m.table
}