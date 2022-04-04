package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BizArticleLookModel = (*customBizArticleLookModel)(nil)

type (
	// BizArticleLookModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBizArticleLookModel.
	BizArticleLookModel interface {
		bizArticleLookModel
	}

	customBizArticleLookModel struct {
		*defaultBizArticleLookModel
	}
)

// NewBizArticleLookModel returns a model for the database table.
func NewBizArticleLookModel(conn sqlx.SqlConn, c cache.CacheConf) BizArticleLookModel {
	return &customBizArticleLookModel{
		defaultBizArticleLookModel: newBizArticleLookModel(conn, c),
	}
}
