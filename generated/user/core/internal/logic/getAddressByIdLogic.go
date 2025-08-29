package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAddressByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAddressByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAddressByIdLogic {
	return &GetAddressByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Gets address by address ID
func (l *GetAddressByIdLogic) GetAddressById(in *core.GetAddressByIdRequest) (*core.GetAddressByIdResponse, error) {
	// todo: add your logic here and delete this line

	return &core.GetAddressByIdResponse{}, nil
}
