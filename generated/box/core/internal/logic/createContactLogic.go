package logic

import (
	"context"

	"core/box/core"
	"core/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateContactLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateContactLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateContactLogic {
	return &CreateContactLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateContactLogic) CreateContact(in *core_core.CreateContactRequest) (*core_core.CreateContactResponse, error) {
	// todo: add your logic here and delete this line

	return &core_core.CreateContactResponse{}, nil
}
