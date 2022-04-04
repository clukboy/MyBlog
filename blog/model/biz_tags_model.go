package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BizTagsModel = (*customBizTagsModel)(nil)

type (
	// BizTagsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBizTagsModel.
	BizTagsModel interface {
		bizTagsModel
	}

	customBizTagsModel struct {
		*defaultBizTagsModel
	}
)

// NewBizTagsModel returns a model for the database table.
func NewBizTagsModel(conn sqlx.SqlConn, c cache.CacheConf) BizTagsModel {
	return &customBizTagsModel{
		defaultBizTagsModel: newBizTagsModel(conn, c),
	}
}
