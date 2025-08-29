package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWishLooksInfoByUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWishLooksInfoByUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWishLooksInfoByUserLogic {
	return &GetWishLooksInfoByUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取用户收藏的look信息
func (l *GetWishLooksInfoByUserLogic) GetWishLooksInfoByUser(in *core.GetWishLooksInfoByUserRequest) (*core.GetWishLooksInfoByUserResponse, error) {
	// todo: add your logic here and delete this line

	return &core.GetWishLooksInfoByUserResponse{}, nil
}
