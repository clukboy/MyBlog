package model

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleCategoryModel = (*customArticleCategoryModel)(nil)

type (
	// ArticleCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleCategoryModel.
	ArticleCategoryModel interface {
		articleCategoryModel
		TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error
	}

	customArticleCategoryModel struct {
		*defaultArticleCategoryModel
	}
)

// NewArticleCategoryModel returns a model for the database table.
func NewArticleCategoryModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleCategoryModel {
	return &customArticleCategoryModel{
		defaultArticleCategoryModel: newArticleCategoryModel(conn, c),
	}
}

func (m *defaultArticleCategoryModel) TransCtx(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, s sqlx.Session) error {
		return fn(ctx, s)
	})
}
