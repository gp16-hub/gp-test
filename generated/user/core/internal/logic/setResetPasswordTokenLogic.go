package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetResetPasswordTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetResetPasswordTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetResetPasswordTokenLogic {
	return &SetResetPasswordTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SetResetPasswordTokenLogic) SetResetPasswordToken(in *core.SetResetPasswordTokenRequest) (*core.SetResetPasswordTokenResponse, error) {
	// todo: add your logic here and delete this line

	return &core.SetResetPasswordTokenResponse{}, nil
}
