package svc

import (
	"mallchat-go/app/user/internal/config"
	"mallchat-go/app/user/internal/service"
	"mallchat-go/app/user/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config     config.Config
	UserModel  model.UserModel
	SmsService *service.SmsService
	Auth       rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DB.DataSource)
	rds := redis.New(c.Redis.Host, func(r *redis.Redis) {
		r.Type = redis.NodeType
		r.Pass = c.Redis.Pass
	})

	return &ServiceContext{
		Config:     c,
		UserModel:  model.NewUserModel(conn),
		SmsService: service.NewSmsService(rds),
	}
}
