package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserByTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserByTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserByTokenLogic {
	return &GetUserByTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserByTokenLogic) GetUserByToken(in *core.GetUserByTokenRequest) (*core.GetUserByTokenResponse, error) {
	// todo: add your logic here and delete this line

	return &core.GetUserByTokenResponse{}, nil
}
