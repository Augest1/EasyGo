package svc

import (
	"EasyGo/services/model/userModel"
	"EasyGo/services/rpc/user/internal/config"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	_ "gorm.io/driver/postgres"
)

type ServiceContext struct {
	Config    config.Config
	UserModel userModel.UsersModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewSqlConn("pgx", c.DataSource)
	return &ServiceContext{
		Config:    c,
		UserModel: userModel.NewUsersModel(conn, c.Cache),
	}
}
