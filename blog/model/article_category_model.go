package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleCategoryModel = (*customArticleCategoryModel)(nil)

type (
	// ArticleCategoryModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleCategoryModel.
	ArticleCategoryModel interface {
		articleCategoryModel
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
