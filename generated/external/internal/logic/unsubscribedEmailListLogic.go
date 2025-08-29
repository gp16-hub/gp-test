package logic

import (
	"context"

	"external/external"
	"external/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type UnsubscribedEmailListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUnsubscribedEmailListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UnsubscribedEmailListLogic {
	return &UnsubscribedEmailListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UnsubscribedEmailListLogic) UnsubscribedEmailList(in *external.UnsubscribedEmailListRequest) (*external.UnsubscribedEmailListResponse, error) {
	// todo: add your logic here and delete this line

	return &external.UnsubscribedEmailListResponse{}, nil
}
