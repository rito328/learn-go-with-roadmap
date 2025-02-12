package logic

import (
	"context"
	"fmt"

	"go_zero_sample/greet/greet"
	"go_zero_sample/greet/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PingLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPingLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PingLogic {
	return &PingLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PingLogic) Ping(in *greet.Request) (*greet.Response, error) {
	fmt.Printf("in %#v\n", in)
	return &greet.Response{
		Pong: "pong",
	}, nil
}
