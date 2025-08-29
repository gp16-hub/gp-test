package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetWishItemVariantsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetWishItemVariantsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetWishItemVariantsLogic {
	return &GetWishItemVariantsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取收藏清单中的所有产品
func (l *GetWishItemVariantsLogic) GetWishItemVariants(in *core.GetWishItemVariantsRequest) (*core.GetWishItemVariantsResponse, error) {
	// todo: add your logic here and delete this line

	return &core.GetWishItemVariantsResponse{}, nil
}
