package logic

import (
	"context"

	"tax/internal/svc"
	"tax/tax"

	"github.com/zeromicro/go-zero/core/logx"
)

type CalculateTaxLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCalculateTaxLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CalculateTaxLogic {
	return &CalculateTaxLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CalculateTaxLogic) CalculateTax(in *tax_tax.TaxRequest) (*tax_tax.TaxResponse, error) {
	// todo: add your logic here and delete this line

	return &tax_tax.TaxResponse{}, nil
}
