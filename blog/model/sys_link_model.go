package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysLinkModel = (*customSysLinkModel)(nil)

type (
	// SysLinkModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysLinkModel.
	SysLinkModel interface {
		sysLinkModel
	}

	customSysLinkModel struct {
		*defaultSysLinkModel
	}
)

// NewSysLinkModel returns a model for the database table.
func NewSysLinkModel(conn sqlx.SqlConn, c cache.CacheConf) SysLinkModel {
	return &customSysLinkModel{
		defaultSysLinkModel: newSysLinkModel(conn, c),
	}
}
