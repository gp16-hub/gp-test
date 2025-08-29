package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchAddWishItemVariantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchAddWishItemVariantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchAddWishItemVariantLogic {
	return &BatchAddWishItemVariantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量添加收藏清单中的产品
func (l *BatchAddWishItemVariantLogic) BatchAddWishItemVariant(in *core.BatchAddWishItemVariantRequest) (*core.BatchAddWishItemVariantResponse, error) {
	// todo: add your logic here and delete this line

	return &core.BatchAddWishItemVariantResponse{}, nil
}
