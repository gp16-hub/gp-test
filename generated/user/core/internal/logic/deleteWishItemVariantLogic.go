package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteWishItemVariantLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteWishItemVariantLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWishItemVariantLogic {
	return &DeleteWishItemVariantLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除收藏清单中的某个产品
func (l *DeleteWishItemVariantLogic) DeleteWishItemVariant(in *core.DeleteWishItemVariantRequest) (*core.DeleteWishItemVariantResponse, error) {
	// todo: add your logic here and delete this line

	return &core.DeleteWishItemVariantResponse{}, nil
}
