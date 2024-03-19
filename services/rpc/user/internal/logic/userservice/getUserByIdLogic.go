package userservicelogic

import (
	"EasyGo/common"
	"EasyGo/services/rpc/user/internal/svc"
	"EasyGo/services/rpc/user/userRpcModel"
	"context"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByIdLogic {
	return &GetUserByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByIdLogic) GetUserById(in *userRpcModel.ReqId) (*userRpcModel.ResUser, error) {
	one, err := l.svcCtx.UserModel.FindOne(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &userRpcModel.ResUser{
		Info: common.Convert.UserS2R(one),
	}, nil
}
