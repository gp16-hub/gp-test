package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserAddressLogic {
	return &CreateUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Create a user address by ID
func (l *CreateUserAddressLogic) CreateUserAddress(in *core.CreateUserAddressRequest) (*core.CreateUserAddressResponse, error) {
	// todo: add your logic here and delete this line

	return &core.CreateUserAddressResponse{}, nil
}
