package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ DiscussModel = (*customDiscussModel)(nil)

type (
	// DiscussModel is an interface to be customized, add more methods here,
	// and implement the added methods in customDiscussModel.
	DiscussModel interface {
		discussModel
	}

	customDiscussModel struct {
		*defaultDiscussModel
	}
)

// NewDiscussModel returns a model for the database table.
func NewDiscussModel(conn sqlx.SqlConn, c cache.CacheConf) DiscussModel {
	return &customDiscussModel{
		defaultDiscussModel: newDiscussModel(conn, c),
	}
}
