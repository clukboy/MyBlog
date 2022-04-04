package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysRoleResourcesModel = (*customSysRoleResourcesModel)(nil)

type (
	// SysRoleResourcesModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysRoleResourcesModel.
	SysRoleResourcesModel interface {
		sysRoleResourcesModel
	}

	customSysRoleResourcesModel struct {
		*defaultSysRoleResourcesModel
	}
)

// NewSysRoleResourcesModel returns a model for the database table.
func NewSysRoleResourcesModel(conn sqlx.SqlConn, c cache.CacheConf) SysRoleResourcesModel {
	return &customSysRoleResourcesModel{
		defaultSysRoleResourcesModel: newSysRoleResourcesModel(conn, c),
	}
}
