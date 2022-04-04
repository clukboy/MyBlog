package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BizAdModel = (*customBizAdModel)(nil)

type (
	// BizAdModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBizAdModel.
	BizAdModel interface {
		bizAdModel
	}

	customBizAdModel struct {
		*defaultBizAdModel
	}
)

// NewBizAdModel returns a model for the database table.
func NewBizAdModel(conn sqlx.SqlConn, c cache.CacheConf) BizAdModel {
	return &customBizAdModel{
		defaultBizAdModel: newBizAdModel(conn, c),
	}
}
