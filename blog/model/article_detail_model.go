package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleDetailModel = (*customArticleDetailModel)(nil)

type (
	// ArticleDetailModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleDetailModel.
	ArticleDetailModel interface {
		articleDetailModel
	}

	customArticleDetailModel struct {
		*defaultArticleDetailModel
	}
)

// NewArticleDetailModel returns a model for the database table.
func NewArticleDetailModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleDetailModel {
	return &customArticleDetailModel{
		defaultArticleDetailModel: newArticleDetailModel(conn, c),
	}
}
