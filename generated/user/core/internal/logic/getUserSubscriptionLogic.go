package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserSubscriptionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserSubscriptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserSubscriptionLogic {
	return &GetUserSubscriptionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserSubscriptionLogic) GetUserSubscription(in *core.GetUserSubscriptionRequest) (*core.GetUserSubscriptionResponse, error) {
	// todo: add your logic here and delete this line

	return &core.GetUserSubscriptionResponse{}, nil
}
