package logic

import (
	"context"

	"external/external"
	"external/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateTaskForCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateTaskForCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateTaskForCustomerLogic {
	return &CreateTaskForCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateTaskForCustomerLogic) CreateTaskForCustomer(in *external.CreateTaskForCustomerRequest) (*external.CreateTaskForCustomerResponse, error) {
	// todo: add your logic here and delete this line

	return &external.CreateTaskForCustomerResponse{}, nil
}
