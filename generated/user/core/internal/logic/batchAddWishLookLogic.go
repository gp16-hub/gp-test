package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type BatchAddWishLookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewBatchAddWishLookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *BatchAddWishLookLogic {
	return &BatchAddWishLookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 批量添加收藏look
func (l *BatchAddWishLookLogic) BatchAddWishLook(in *core.BatchAddWishLookRequest) (*core.BatchAddWishLookResponse, error) {
	// todo: add your logic here and delete this line

	return &core.BatchAddWishLookResponse{}, nil
}
