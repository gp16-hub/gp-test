package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindOrCreateTermsOfUseLogByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindOrCreateTermsOfUseLogByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindOrCreateTermsOfUseLogByUserLogic {
	return &FindOrCreateTermsOfUseLogByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户接受的使用协议的日志
func (l *FindOrCreateTermsOfUseLogByUserLogic) FindOrCreateTermsOfUseLogByUser(in *core.FindOrCreateTermsOfUseLogByUserRequest) (*core.FindOrCreateTermsOfUseLogByUserResponse, error) {
	// todo: add your logic here and delete this line

	return &core.FindOrCreateTermsOfUseLogByUserResponse{}, nil
}
