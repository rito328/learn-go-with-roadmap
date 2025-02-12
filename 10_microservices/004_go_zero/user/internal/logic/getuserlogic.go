package logic

import (
	"context"
	"fmt"

	"go_zero_sample/user/internal/svc"
	"go_zero_sample/user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserLogic {
	return &GetUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserLogic) GetUser(req *types.GetUserRequest) (resp *types.GetUserResponse, err error) {
	// ログを記録
	l.Logger.Infof("Received request: %v", req)

	resp = &types.GetUserResponse{
		Message: fmt.Sprintf("Hello World, UserId: %s", req.UserId),
	}

	// ログを記録
	l.Logger.Infof("Response: %v", resp)

	return resp, nil
}
