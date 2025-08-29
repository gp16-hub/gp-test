package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByQueryLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByQueryLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByQueryLogic {
	return &GetUserByQueryLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByQueryLogic) GetUserByQuery(in *core.GetUserByQueryRequest) (*core.GetUserByQueryResponse, error) {
	// todo: add your logic here and delete this line

	return &core.GetUserByQueryResponse{}, nil
}
