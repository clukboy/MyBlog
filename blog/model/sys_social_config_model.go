package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysSocialConfigModel = (*customSysSocialConfigModel)(nil)

type (
	// SysSocialConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysSocialConfigModel.
	SysSocialConfigModel interface {
		sysSocialConfigModel
	}

	customSysSocialConfigModel struct {
		*defaultSysSocialConfigModel
	}
)

// NewSysSocialConfigModel returns a model for the database table.
func NewSysSocialConfigModel(conn sqlx.SqlConn, c cache.CacheConf) SysSocialConfigModel {
	return &customSysSocialConfigModel{
		defaultSysSocialConfigModel: newSysSocialConfigModel(conn, c),
	}
}
