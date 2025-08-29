package logic

import (
	"context"

	"core/internal/svc"
	"core/promotion"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListPromotionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListPromotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListPromotionLogic {
	return &ListPromotionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// promotion
func (l *ListPromotionLogic) ListPromotion(in *promotion.ListPromotionReq) (*promotion.ListPromotionRsp, error) {
	// todo: add your logic here and delete this line

	return &promotion.ListPromotionRsp{}, nil
}
