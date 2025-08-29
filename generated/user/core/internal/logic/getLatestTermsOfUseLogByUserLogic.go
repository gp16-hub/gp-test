package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetLatestTermsOfUseLogByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetLatestTermsOfUseLogByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetLatestTermsOfUseLogByUserLogic {
	return &GetLatestTermsOfUseLogByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户最新接受的使用协议日志
func (l *GetLatestTermsOfUseLogByUserLogic) GetLatestTermsOfUseLogByUser(in *core.GetLatestTermsOfUseLogByUserRequest) (*core.GetLatestTermsOfUseLogByUserResponse, error) {
	// todo: add your logic here and delete this line

	return &core.GetLatestTermsOfUseLogByUserResponse{}, nil
}
