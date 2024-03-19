// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package userservice

import (
	"context"

	"EasyGo/services/rpc/user/userRpcModel"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	ReqCreateUser = userRpcModel.ReqCreateUser
	ReqId         = userRpcModel.ReqId
	ResUser       = userRpcModel.ResUser

	UserService interface {
		CreateUser(ctx context.Context, in *ReqCreateUser, opts ...grpc.CallOption) (*ResUser, error)
		GetUserById(ctx context.Context, in *ReqId, opts ...grpc.CallOption) (*ResUser, error)
		UpdateUser(ctx context.Context, in *ReqCreateUser, opts ...grpc.CallOption) (*ResUser, error)
	}

	defaultUserService struct {
		cli zrpc.Client
	}
)

func NewUserService(cli zrpc.Client) UserService {
	return &defaultUserService{
		cli: cli,
	}
}

func (m *defaultUserService) CreateUser(ctx context.Context, in *ReqCreateUser, opts ...grpc.CallOption) (*ResUser, error) {
	client := userRpcModel.NewUserServiceClient(m.cli.Conn())
	return client.CreateUser(ctx, in, opts...)
}

func (m *defaultUserService) GetUserById(ctx context.Context, in *ReqId, opts ...grpc.CallOption) (*ResUser, error) {
	client := userRpcModel.NewUserServiceClient(m.cli.Conn())
	return client.GetUserById(ctx, in, opts...)
}

func (m *defaultUserService) UpdateUser(ctx context.Context, in *ReqCreateUser, opts ...grpc.CallOption) (*ResUser, error) {
	client := userRpcModel.NewUserServiceClient(m.cli.Conn())
	return client.UpdateUser(ctx, in, opts...)
}
