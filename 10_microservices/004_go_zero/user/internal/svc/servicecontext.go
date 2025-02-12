package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"go_zero_sample/user/internal/config"
	"go_zero_sample/user/internal/middleware"
)

type ServiceContext struct {
	Config                          config.Config
	RequestDurationLoggerMiddleware rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                          c,
		RequestDurationLoggerMiddleware: middleware.NewRequestDurationLoggerMiddleware().Handle,
	}
}
