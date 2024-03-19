// Code generated by goctl. DO NOT EDIT.

package userModel

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	usersFieldNames          = builder.RawFieldNames(&Users{}, true)
	usersRows                = strings.Join(usersFieldNames, ",")
	usersRowsExpectAutoSet   = strings.Join(stringx.Remove(usersFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"), ",")
	usersRowsWithPlaceHolder = builder.PostgreSqlJoin(stringx.Remove(usersFieldNames, "id", "create_at", "create_time", "created_at", "update_at", "update_time", "updated_at"))

	cachePublicUsersIdPrefix       = "cache:public:users:id:"
	cachePublicUsersUsernamePrefix = "cache:public:users:username:"
)

type (
	usersModel interface {
		Insert(ctx context.Context, data *Users) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Users, error)
		FindOneByUsername(ctx context.Context, username string) (*Users, error)
		Update(ctx context.Context, data *Users) error
		Delete(ctx context.Context, id int64) error
	}

	defaultUsersModel struct {
		sqlc.CachedConn
		table string
	}

	Users struct {
		Id        int64          `db:"id"`
		Username  string         `db:"username"`
		Nickname  string         `db:"nickname"`
		Name      sql.NullString `db:"name"`
		CreatedAt time.Time      `db:"created_at"`
		UpdatedAt time.Time      `db:"updated_at"`
		DeletedAt sql.NullTime   `db:"deleted_at"`
		Status    int64          `db:"status"`
		Signature sql.NullString `db:"signature"`
		Avatar    string         `db:"avatar"`
		Sex       int64          `db:"sex"`
		Mobile    string         `db:"mobile"`
		IdCardNum sql.NullString `db:"id_card_num"`
		Birth     sql.NullTime   `db:"birth"`
	}
)

func newUsersModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultUsersModel {
	return &defaultUsersModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      `"public"."users"`,
	}
}

func (m *defaultUsersModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, id)
	publicUsersUsernameKey := fmt.Sprintf("%s%v", cachePublicUsersUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where id = $1", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, publicUsersIdKey, publicUsersUsernameKey)
	return err
}

func (m *defaultUsersModel) FindOne(ctx context.Context, id int64) (*Users, error) {
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, id)
	var resp Users
	err := m.QueryRowCtx(ctx, &resp, publicUsersIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
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

func (m *defaultUsersModel) FindOneByUsername(ctx context.Context, username string) (*Users, error) {
	publicUsersUsernameKey := fmt.Sprintf("%s%v", cachePublicUsersUsernamePrefix, username)
	var resp Users
	err := m.QueryRowIndexCtx(ctx, &resp, publicUsersUsernameKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where username = $1 limit 1", usersRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, username); err != nil {
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

func (m *defaultUsersModel) Insert(ctx context.Context, data *Users) (sql.Result, error) {
	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, data.Id)
	publicUsersUsernameKey := fmt.Sprintf("%s%v", cachePublicUsersUsernamePrefix, data.Username)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)", m.table, usersRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Username, data.Nickname, data.Name, data.DeletedAt, data.Status, data.Signature, data.Avatar, data.Sex, data.Mobile, data.IdCardNum, data.Birth)
	}, publicUsersIdKey, publicUsersUsernameKey)
	return ret, err
}

func (m *defaultUsersModel) Update(ctx context.Context, newData *Users) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	publicUsersIdKey := fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, data.Id)
	publicUsersUsernameKey := fmt.Sprintf("%s%v", cachePublicUsersUsernamePrefix, data.Username)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where id = $1", m.table, usersRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.Id, newData.Username, newData.Nickname, newData.Name, newData.DeletedAt, newData.Status, newData.Signature, newData.Avatar, newData.Sex, newData.Mobile, newData.IdCardNum, newData.Birth)
	}, publicUsersIdKey, publicUsersUsernameKey)
	return err
}

func (m *defaultUsersModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cachePublicUsersIdPrefix, primary)
}

func (m *defaultUsersModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where id = $1 limit 1", usersRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUsersModel) tableName() string {
	return m.table
}