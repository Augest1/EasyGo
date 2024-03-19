package userservicelogic

import (
	"EasyGo/services/model/usermodel"
	"EasyGo/services/rpc/user/internal/svc"
	"EasyGo/services/rpc/user/userRpcModel"
	"context"
	"database/sql"
	"time"

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
	var birth, deletedAt time.Time
	var err error
	if in.Info.Birth != "" {
		birth, err = time.Parse(time.DateTime, in.Info.Birth)
		if err != nil {
			logx.Info(err)
			return nil, err
		}
	}

	if in.Info.DeletedAt != "" {
		deletedAt, err = time.Parse(time.DateTime, in.Info.DeletedAt)
		if err != nil {
			logx.Info(err)
			return nil, err
		}
	}
	err = l.svcCtx.UserModel.Update(l.ctx, &usermodel.Users{
		Username: in.Info.Username,
		Nickname: in.Info.Nickname,
		Name: sql.NullString{
			String: in.Info.Name,
			Valid:  in.Info.Name != "",
		},
		DeletedAt: sql.NullTime{
			Time:  deletedAt,
			Valid: in.Info.DeletedAt != "",
		},
		Status: in.Info.Status,
		Signature: sql.NullString{
			String: in.Info.Signature,
			Valid:  in.Info.Signature != "",
		},
		Avatar: in.Info.Avatar,
		Sex:    in.Info.Sex,
		Mobile: in.Info.Mobile,
		IdCardNum: sql.NullString{
			String: in.Info.IdCardNum,
			Valid:  in.Info.IdCardNum != "",
		},
		Birth: sql.NullTime{
			Time:  birth,
			Valid: in.Info.Birth != "",
		},
	})
	if err != nil {
		return nil, err
	}
	return &userRpcModel.ResUser{
		Info: in.Info,
	}, nil
}
