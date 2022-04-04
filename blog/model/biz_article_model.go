package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BizArticleModel = (*customBizArticleModel)(nil)

type (
	// BizArticleModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBizArticleModel.
	BizArticleModel interface {
		bizArticleModel
	}

	customBizArticleModel struct {
		*defaultBizArticleModel
	}
)

// NewBizArticleModel returns a model for the database table.
func NewBizArticleModel(conn sqlx.SqlConn, c cache.CacheConf) BizArticleModel {
	return &customBizArticleModel{
		defaultBizArticleModel: newBizArticleModel(conn, c),
	}
}
