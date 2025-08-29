package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserAddressLogic {
	return &UpdateUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Update a user address
func (l *UpdateUserAddressLogic) UpdateUserAddress(in *core.UpdateUserAddressRequest) (*core.UpdateUserAddressResponse, error) {
	// todo: add your logic here and delete this line

	return &core.UpdateUserAddressResponse{}, nil
}
