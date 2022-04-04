package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ BizCommentModel = (*customBizCommentModel)(nil)

type (
	// BizCommentModel is an interface to be customized, add more methods here,
	// and implement the added methods in customBizCommentModel.
	BizCommentModel interface {
		bizCommentModel
	}

	customBizCommentModel struct {
		*defaultBizCommentModel
	}
)

// NewBizCommentModel returns a model for the database table.
func NewBizCommentModel(conn sqlx.SqlConn, c cache.CacheConf) BizCommentModel {
	return &customBizCommentModel{
		defaultBizCommentModel: newBizCommentModel(conn, c),
	}
}
