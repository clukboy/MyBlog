package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysUpdateRecordeModel = (*customSysUpdateRecordeModel)(nil)

type (
	// SysUpdateRecordeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysUpdateRecordeModel.
	SysUpdateRecordeModel interface {
		sysUpdateRecordeModel
	}

	customSysUpdateRecordeModel struct {
		*defaultSysUpdateRecordeModel
	}
)

// NewSysUpdateRecordeModel returns a model for the database table.
func NewSysUpdateRecordeModel(conn sqlx.SqlConn, c cache.CacheConf) SysUpdateRecordeModel {
	return &customSysUpdateRecordeModel{
		defaultSysUpdateRecordeModel: newSysUpdateRecordeModel(conn, c),
	}
}
