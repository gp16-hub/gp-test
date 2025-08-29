package logic

import (
	"context"

	"external/external"
	"external/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateSubscriptionListForCustomerLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateSubscriptionListForCustomerLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateSubscriptionListForCustomerLogic {
	return &CreateSubscriptionListForCustomerLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateSubscriptionListForCustomerLogic) CreateSubscriptionListForCustomer(in *external.CreateSubscriptionListForCustomerRequest) (*external.CreateSubscriptionListForCustomerResponse, error) {
	// todo: add your logic here and delete this line

	return &external.CreateSubscriptionListForCustomerResponse{}, nil
}
