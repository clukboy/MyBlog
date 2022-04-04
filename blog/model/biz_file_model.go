package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BizFileModel = (*customBizFileModel)(nil)

type (
	// BizFileModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBizFileModel.
	BizFileModel interface {
		bizFileModel
	}

	customBizFileModel struct {
		*defaultBizFileModel
	}
)

// NewBizFileModel returns a model for the database table.
func NewBizFileModel(conn sqlx.SqlConn, c cache.CacheConf) BizFileModel {
	return &customBizFileModel{
		defaultBizFileModel: newBizFileModel(conn, c),
	}
}
