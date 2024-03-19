package userservicelogic

import (
	"EasyGo/common"
	"EasyGo/services/rpc/user/internal/svc"
	"EasyGo/services/rpc/user/userRpcModel"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserLogic) CreateUser(in *userRpcModel.ReqCreateUser) (*userRpcModel.ResUser, error) {
	one, err := l.svcCtx.UserModel.Insert(l.ctx, common.Convert.UserR2S(in.Info))
	if err != nil {
		return nil, err
	}
	in.Info.Id, _ = one.LastInsertId()
	return &userRpcModel.ResUser{
		Info: in.Info,
	}, nil
}
