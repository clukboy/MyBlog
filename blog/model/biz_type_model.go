package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BizTypeModel = (*customBizTypeModel)(nil)

type (
	// BizTypeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBizTypeModel.
	BizTypeModel interface {
		bizTypeModel
	}

	customBizTypeModel struct {
		*defaultBizTypeModel
	}
)

// NewBizTypeModel returns a model for the database table.
func NewBizTypeModel(conn sqlx.SqlConn, c cache.CacheConf) BizTypeModel {
	return &customBizTypeModel{
		defaultBizTypeModel: newBizTypeModel(conn, c),
	}
}
