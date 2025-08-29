package logic

import (
	"context"

	"core/internal/svc"
	"core/promotion"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddPromotionToEmailsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddPromotionToEmailsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddPromotionToEmailsLogic {
	return &AddPromotionToEmailsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddPromotionToEmailsLogic) AddPromotionToEmails(in *promotion.AddPromotionToEmailsReq) (*promotion.AddPromotionToEmailsRsp, error) {
	// todo: add your logic here and delete this line

	return &promotion.AddPromotionToEmailsRsp{}, nil
}
