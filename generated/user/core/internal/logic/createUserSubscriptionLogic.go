package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateUserSubscriptionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateUserSubscriptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserSubscriptionLogic {
	return &CreateUserSubscriptionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateUserSubscriptionLogic) CreateUserSubscription(in *core.CreateUserSubscriptionRequest) (*core.CreateUserSubscriptionResponse, error) {
	// todo: add your logic here and delete this line

	return &core.CreateUserSubscriptionResponse{}, nil
}
