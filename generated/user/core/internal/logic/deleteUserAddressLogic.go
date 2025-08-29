package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteUserAddressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteUserAddressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteUserAddressLogic {
	return &DeleteUserAddressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Delete a user address by ID
func (l *DeleteUserAddressLogic) DeleteUserAddress(in *core.DeleteUserAddressRequest) (*core.DeleteUserAddressResponse, error) {
	// todo: add your logic here and delete this line

	return &core.DeleteUserAddressResponse{}, nil
}
