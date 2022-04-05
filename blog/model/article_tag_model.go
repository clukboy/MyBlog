package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ArticleTagModel = (*customArticleTagModel)(nil)

type (
	// ArticleTagModel is an interface to be customized, add more methods here,
	// and implement the added methods in customArticleTagModel.
	ArticleTagModel interface {
		articleTagModel
	}

	customArticleTagModel struct {
		*defaultArticleTagModel
	}
)

// NewArticleTagModel returns a model for the database table.
func NewArticleTagModel(conn sqlx.SqlConn, c cache.CacheConf) ArticleTagModel {
	return &customArticleTagModel{
		defaultArticleTagModel: newArticleTagModel(conn, c),
	}
}
