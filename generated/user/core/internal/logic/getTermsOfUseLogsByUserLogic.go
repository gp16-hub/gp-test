package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTermsOfUseLogsByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTermsOfUseLogsByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTermsOfUseLogsByUserLogic {
	return &GetTermsOfUseLogsByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户接受的使用协议的日志
func (l *GetTermsOfUseLogsByUserLogic) GetTermsOfUseLogsByUser(in *core.GetTermsOfUseLogsByUserRequest) (*core.GetTermsOfUseLogsByUserResponse, error) {
	// todo: add your logic here and delete this line

	return &core.GetTermsOfUseLogsByUserResponse{}, nil
}
