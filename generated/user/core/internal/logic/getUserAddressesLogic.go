package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAddressesLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAddressesLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAddressesLogic {
	return &GetUserAddressesLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Gets a user address by ID
func (l *GetUserAddressesLogic) GetUserAddresses(in *core.GetUserAddressesRequest) (*core.GetUserAddressesResponse, error) {
	// todo: add your logic here and delete this line

	return &core.GetUserAddressesResponse{}, nil
}
