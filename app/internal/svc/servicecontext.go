package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"mallchat-go/app/internal/config"
	"mallchat-go/app/internal/middleware"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Auth:   middleware.NewAuthMiddleware().Handle,
	}
}
