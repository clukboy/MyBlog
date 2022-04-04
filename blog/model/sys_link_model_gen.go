// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	sysLinkFieldNames          = builder.RawFieldNames(&SysLink{}, true)
	sysLinkRows                = strings.Join(sysLinkFieldNames, ",")
	sysLinkRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(sysLinkFieldNames, "id", "create_time"))

	cachePublicSysLinkIdPrefix = "cache:public:sysLink:id:"
)

type (
	sysLinkModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *SysLink) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*SysLink, error)
		Update(ctx context.Context, session sqlx.Session, data *SysLink) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultSysLinkModel struct {
		sqlc.CachedConn
		table string
	}

	SysLink struct {
		Id              int64     `db:"id"`
		Url             string    `db:"url"`         // 链接地址
		Name            string    `db:"name"`        // 链接名
		Description     string    `db:"description"` // 链接介绍
		Email           string    `db:"email"`       // 友链站长邮箱
		Qq              string    `db:"qq"`          // 友链站长QQ
		Favicon         string    `db:"favicon"`
		Status          int64     `db:"status"`            // 状态
		HomePageDisplay int64     `db:"home_page_display"` // 是否首页显示
		Remark          string    `db:"remark"`            // 备注
		Source          string    `db:"source"`            // 来源：管理员添加、自动申请
		CreateTime      time.Time `db:"create_time"`       // 添加时间
		UpdateTime      time.Time `db:"update_time"`       // 更新时间
	}
)

func newSysLinkModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultSysLinkModel {
	return &defaultSysLinkModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."sys_link"`,
	}
}

func (m *defaultSysLinkModel) Insert(ctx context.Context, session sqlx.Session, data *SysLink) (sql.Result, error) {
	query, args, err := sq.Insert(m.table).Columns(sysLinkRows).Values(data.Id, data.Url, data.Name, data.Description, data.Email, data.Qq, data.Favicon, data.Status, data.HomePageDisplay, data.Remark, data.Source, data.CreateTime, data.UpdateTime).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	publicSysLinkIdKey := fmt.Sprintf("%s%v", cachePublicSysLinkIdPrefix, data.Id)

	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicSysLinkIdKey)
	return ret, err
}

func (m *defaultSysLinkModel) FindOne(ctx context.Context, id int64) (*SysLink, error) {
	publicSysLinkIdKey := fmt.Sprintf("%s%v", cachePublicSysLinkIdPrefix, id)
	var resp SysLink
	err := m.QueryRowCtx(ctx, &resp, publicSysLinkIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", sysLinkRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultSysLinkModel) Update(ctx context.Context, session sqlx.Session, data *SysLink) error {
	publicSysLinkIdKey := fmt.Sprintf("%s%v", cachePublicSysLinkIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, sysLinkRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.Url, data.Name, data.Description, data.Email, data.Qq, data.Favicon, data.Status, data.HomePageDisplay, data.Remark, data.Source, time.Now())

		}
		return conn.ExecCtx(ctx, query, data.Id, data.Url, data.Name, data.Description, data.Email, data.Qq, data.Favicon, data.Status, data.HomePageDisplay, data.Remark, data.Source, time.Now())
	}, publicSysLinkIdKey)
	return err
}

func (m *defaultSysLinkModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query, args, err := sq.Delete(m.table).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	publicSysLinkIdKey := fmt.Sprintf("%s%v", cachePublicSysLinkIdPrefix, id)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicSysLinkIdKey)
	return err
}

func (m *defaultSysLinkModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicSysLinkIdPrefix, primary)
}

func (m *defaultSysLinkModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", sysLinkRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultSysLinkModel) tableName() string {
	return m.table
}