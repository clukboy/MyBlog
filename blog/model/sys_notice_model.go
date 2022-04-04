package model

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ SysNoticeModel = (*customSysNoticeModel)(nil)

type (
	// SysNoticeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customSysNoticeModel.
	SysNoticeModel interface {
		sysNoticeModel
	}

	customSysNoticeModel struct {
		*defaultSysNoticeModel
	}
)

// NewSysNoticeModel returns a model for the database table.
func NewSysNoticeModel(conn sqlx.SqlConn, c cache.CacheConf) SysNoticeModel {
	return &customSysNoticeModel{
		defaultSysNoticeModel: newSysNoticeModel(conn, c),
	}
}
