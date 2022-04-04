package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysTemplateModel = (*customSysTemplateModel)(nil)

type (
	// SysTemplateModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysTemplateModel.
	SysTemplateModel interface {
		sysTemplateModel
	}

	customSysTemplateModel struct {
		*defaultSysTemplateModel
	}
)

// NewSysTemplateModel returns a model for the database table.
func NewSysTemplateModel(conn sqlx.SqlConn, c cache.CacheConf) SysTemplateModel {
	return &customSysTemplateModel{
		defaultSysTemplateModel: newSysTemplateModel(conn, c),
	}
}
