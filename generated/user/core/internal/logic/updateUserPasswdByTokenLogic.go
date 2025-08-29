package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserPasswdByTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserPasswdByTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserPasswdByTokenLogic {
	return &UpdateUserPasswdByTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateUserPasswdByTokenLogic) UpdateUserPasswdByToken(in *core.UpdateUserPasswdByTokenRequest) (*core.UpdateUserPasswdByTokenResponse, error) {
	// todo: add your logic here and delete this line

	return &core.UpdateUserPasswdByTokenResponse{}, nil
}
