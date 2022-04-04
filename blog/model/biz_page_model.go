package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BizPageModel = (*customBizPageModel)(nil)

type (
	// BizPageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBizPageModel.
	BizPageModel interface {
		bizPageModel
	}

	customBizPageModel struct {
		*defaultBizPageModel
	}
)

// NewBizPageModel returns a model for the database table.
func NewBizPageModel(conn sqlx.SqlConn, c cache.CacheConf) BizPageModel {
	return &customBizPageModel{
		defaultBizPageModel: newBizPageModel(conn, c),
	}
}
