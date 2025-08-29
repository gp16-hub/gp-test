package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthOrCreateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAuthOrCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AuthOrCreateUserLogic {
	return &AuthOrCreateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AuthOrCreateUserLogic) AuthOrCreateUser(in *core.AuthOrCreateUserRequest) (*core.AuthOrCreateUserResponse, error) {
	// todo: add your logic here and delete this line

	return &core.AuthOrCreateUserResponse{}, nil
}
