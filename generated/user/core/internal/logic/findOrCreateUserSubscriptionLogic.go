package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOrCreateUserSubscriptionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOrCreateUserSubscriptionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOrCreateUserSubscriptionLogic {
	return &FindOrCreateUserSubscriptionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindOrCreateUserSubscriptionLogic) FindOrCreateUserSubscription(in *core.FindOrCreateUserSubscriptionRequest) (*core.FindOrCreateUserSubscriptionResponse, error) {
	// todo: add your logic here and delete this line

	return &core.FindOrCreateUserSubscriptionResponse{}, nil
}
