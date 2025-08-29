package logic

import (
	"context"

	"core/internal/svc"
	"core/user/core"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetBirthdayUsersOfCurrentMonthLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetBirthdayUsersOfCurrentMonthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetBirthdayUsersOfCurrentMonthLogic {
	return &GetBirthdayUsersOfCurrentMonthLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Get birthday users of current month
func (l *GetBirthdayUsersOfCurrentMonthLogic) GetBirthdayUsersOfCurrentMonth(in *core.GetBirthdayUsersOfCurrentMonthRequest) (*core.GetBirthdayUsersOfCurrentMonthResponse, error) {
	// todo: add your logic here and delete this line

	return &core.GetBirthdayUsersOfCurrentMonthResponse{}, nil
}
