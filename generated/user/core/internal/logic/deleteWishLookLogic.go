package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteWishLookLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteWishLookLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteWishLookLogic {
	return &DeleteWishLookLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除收藏look
func (l *DeleteWishLookLogic) DeleteWishLook(in *core.DeleteWishLookRequest) (*core.DeleteWishLookResponse, error) {
	// todo: add your logic here and delete this line

	return &core.DeleteWishLookResponse{}, nil
}
