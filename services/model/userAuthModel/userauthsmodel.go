package userAuthModel

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserAuthsModel = (*customUserAuthsModel)(nil)

type (
	// UserAuthsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserAuthsModel.
	UserAuthsModel interface {
		userAuthsModel
	}

	customUserAuthsModel struct {
		*defaultUserAuthsModel
	}
)

// NewUserAuthsModel returns a model for the database table.
func NewUserAuthsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserAuthsModel {
	return &customUserAuthsModel{
		defaultUserAuthsModel: newUserAuthsModel(conn, c, opts...),
	}
}
