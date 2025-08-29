package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Gets a user by ID
func (l *GetUserLogic) GetUser(in *core.GetUserRequest) (*core.GetUserResponse, error) {
	// todo: add your logic here and delete this line

	return &core.GetUserResponse{}, nil
}
