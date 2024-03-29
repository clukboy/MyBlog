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
	usersFieldNames          = builder.RawFieldNames(&Users{}, true)
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(usersFieldNames, "id", "create_time"))

	cachePublicUsersIdPrefix       = "cache:public:users:id:"
	cachePublicUsersEmailPrefix    = "cache:public:users:email:"
	cachePublicUsersPhonePrefix    = "cache:public:users:phone:"
	cachePublicUsersUserNamePrefix = "cache:public:users:userName:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		FindOneByEmail(ctx context.Context, email string) (*Users, error)
		FindOneByPhone(ctx context.Context, phone string) (*Users, error)
		FindOneByUserName(ctx context.Context, userName string) (*Users, error)
		Update(ctx context.Context, session sqlx.Session, data *Users) error
		Delete(ctx context.Context, session sqlx.Session, id int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id         int64     `db:"id"`          // 用户id
		UserName   string    `db:"user_name"`   // 用户名
		NickName   string    `db:"nick_name"`   // 用户昵称
		Password   string    `db:"password"`    // 密码
		Email      string    `db:"email"`       // 邮箱
		Avatar     string    `db:"avatar"`      // 头像
		Birthday   time.Time `db:"birthday"`    // 生日
		Age        int64     `db:"age"`         // 年龄
		Phone      string    `db:"phone"`       // 手机号
		CreateTime time.Time `db:"create_time"` // 创建时间
		UpdateTime time.Time `db:"update_time"` // 修改时间
	}
)

func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      `"public"."users"`,
	}
}

func (m *defaultUsersModel) Insert(ctx context.Context, session sqlx.Session, data *Users) (sql.Result, error) {
	query, args, err := sq.Insert(m.table).Columns(usersRows).Values(data.Id, data.UserName, data.NickName, data.Password, data.Email, data.Avatar, data.Birthday, data.Age, data.Phone, data.CreateTime, data.UpdateTime).
		PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return nil, err
	}
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, data.Id)
	publicUsersEmailKey := fmt.Sprintf("%s%v", cachePublicUsersEmailPrefix, data.Email)
	publicUsersPhoneKey := fmt.Sprintf("%s%v", cachePublicUsersPhonePrefix, data.Phone)
	publicUsersUserNameKey := fmt.Sprintf("%s%v", cachePublicUsersUserNamePrefix, data.UserName)

	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicUsersIdKey, publicUsersEmailKey, publicUsersPhoneKey, publicUsersUserNameKey)
	return ret, err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, id)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, publicUsersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where id = $1 limit 1", usersRows, m.table)
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

func (m *defaultUsersModel) FindOneByEmail(ctx context.Context, email string) (*Users, error) {
	publicUsersEmailKey := fmt.Sprintf("%s%v", cachePublicUsersEmailPrefix, email)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, publicUsersEmailKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where email = $1 limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, email); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByPhone(ctx context.Context, phone string) (*Users, error) {
	publicUsersPhoneKey := fmt.Sprintf("%s%v", cachePublicUsersPhonePrefix, phone)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, publicUsersPhoneKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where phone = $1 limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, phone); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) FindOneByUserName(ctx context.Context, userName string) (*Users, error) {
	publicUsersUserNameKey := fmt.Sprintf("%s%v", cachePublicUsersUserNamePrefix, userName)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, publicUsersUserNameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) (i interface{}, e error) {
		query := fmt.Sprintf("select %s from %s where user_name = $1 limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, userName); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUsersModel) Update(ctx context.Context, session sqlx.Session, data *Users) error {
	publicUsersPhoneKey := fmt.Sprintf("%s%v", cachePublicUsersPhonePrefix, data.Phone)
	publicUsersUserNameKey := fmt.Sprintf("%s%v", cachePublicUsersUserNamePrefix, data.UserName)
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, data.Id)
	publicUsersEmailKey := fmt.Sprintf("%s%v", cachePublicUsersEmailPrefix, data.Email)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, usersRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.UserName, data.NickName, data.Password, data.Email, data.Avatar, data.Birthday, data.Age, data.Phone, time.Now())

		}
		return conn.ExecCtx(ctx, query, data.Id, data.UserName, data.NickName, data.Password, data.Email, data.Avatar, data.Birthday, data.Age, data.Phone, time.Now())
	}, publicUsersIdKey, publicUsersEmailKey, publicUsersPhoneKey, publicUsersUserNameKey)
	return err
}

func (m *defaultUsersModel) Delete(ctx context.Context, session sqlx.Session, id int64) error {
	query, args, err := sq.Delete(m.table).Where(sq.Eq{"id": id}).PlaceholderFormat(sq.Dollar).ToSql()
	if err != nil {
		return err
	}
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, id)
	publicUsersEmailKey := fmt.Sprintf("%s%v", cachePublicUsersEmailPrefix, data.Email)
	publicUsersPhoneKey := fmt.Sprintf("%s%v", cachePublicUsersPhonePrefix, data.Phone)
	publicUsersUserNameKey := fmt.Sprintf("%s%v", cachePublicUsersUserNamePrefix, data.UserName)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		if session != nil {
			return session.ExecCtx(ctx, query, args...)
		}
		return conn.ExecCtx(ctx, query, args...)
	}, publicUsersPhoneKey, publicUsersUserNameKey, publicUsersIdKey, publicUsersEmailKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}
