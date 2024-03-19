package svc

import (
	"EasyGo/services/model/usermodel"
	"EasyGo/services/rpc/user/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	_ "gorm.io/driver/postgres"
)

type ServiceContext struct {
	Config    config.Config
	UserModel usermodel.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("pgx", c.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: usermodel.NewUsersModel(conn, c.Cache),
	}
}
