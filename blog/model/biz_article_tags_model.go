package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BizArticleTagsModel = (*customBizArticleTagsModel)(nil)

type (
	// BizArticleTagsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBizArticleTagsModel.
	BizArticleTagsModel interface {
		bizArticleTagsModel
	}

	customBizArticleTagsModel struct {
		*defaultBizArticleTagsModel
	}
)

// NewBizArticleTagsModel returns a model for the database table.
func NewBizArticleTagsModel(conn sqlx.SqlConn, c cache.CacheConf) BizArticleTagsModel {
	return &customBizArticleTagsModel{
		defaultBizArticleTagsModel: newBizArticleTagsModel(conn, c),
	}
}
