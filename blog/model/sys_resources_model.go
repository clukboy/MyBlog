package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysResourcesModel = (*customSysResourcesModel)(nil)

type (
	// SysResourcesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysResourcesModel.
	SysResourcesModel interface {
		sysResourcesModel
	}

	customSysResourcesModel struct {
		*defaultSysResourcesModel
	}
)

// NewSysResourcesModel returns a model for the database table.
func NewSysResourcesModel(conn sqlx.SqlConn, c cache.CacheConf) SysResourcesModel {
	return &customSysResourcesModel{
		defaultSysResourcesModel: newSysResourcesModel(conn, c),
	}
}
