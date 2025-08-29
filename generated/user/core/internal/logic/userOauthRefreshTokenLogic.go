package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserOauthRefreshTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserOauthRefreshTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserOauthRefreshTokenLogic {
	return &UserOauthRefreshTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserOauthRefreshTokenLogic) UserOauthRefreshToken(in *core.UserOauthRefreshTokenRequest) (*core.UserOauthRefreshTokenResponse, error) {
	// todo: add your logic here and delete this line

	return &core.UserOauthRefreshTokenResponse{}, nil
}
