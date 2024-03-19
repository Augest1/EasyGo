package userservicelogic

import (
	"EasyGo/common"
	"EasyGo/services/rpc/user/internal/svc"
	"EasyGo/services/rpc/user/userRpcModel"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserLogic) UpdateUser(in *userRpcModel.ReqCreateUser) (*userRpcModel.ResUser, error) {
	err := l.svcCtx.UserModel.Update(l.ctx, common.Convert.UserR2S(in.Info))
	if err != nil {
		return nil, err
	}
	return &userRpcModel.ResUser{
		Info: in.Info,
	}, nil
}
