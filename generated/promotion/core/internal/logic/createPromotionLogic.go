package logic

import (
	"context"

	"core/internal/svc"
	"core/promotion"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePromotionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePromotionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePromotionLogic {
	return &CreatePromotionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreatePromotionLogic) CreatePromotion(in *promotion.CreatePromotionReq) (*promotion.CreatePromotionRsp, error) {
	// todo: add your logic here and delete this line

	return &promotion.CreatePromotionRsp{}, nil
}
