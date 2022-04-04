package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BizArticleLoveModel = (*customBizArticleLoveModel)(nil)

type (
	// BizArticleLoveModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBizArticleLoveModel.
	BizArticleLoveModel interface {
		bizArticleLoveModel
	}

	customBizArticleLoveModel struct {
		*defaultBizArticleLoveModel
	}
)

// NewBizArticleLoveModel returns a model for the database table.
func NewBizArticleLoveModel(conn sqlx.SqlConn, c cache.CacheConf) BizArticleLoveModel {
	return &customBizArticleLoveModel{
		defaultBizArticleLoveModel: newBizArticleLoveModel(conn, c),
	}
}
