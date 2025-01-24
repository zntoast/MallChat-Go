package svc

import (
	"context"
	"mallchat-go/app/internal/config"
	"mallchat-go/app/internal/middleware"
	"mallchat-go/app/internal/model"
	"mallchat-go/app/internal/pkg/utils"

	"github.com/importcjj/sensitive"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config   config.Config
	Auth     rest.Middleware
	RedisCli *utils.RedisClient
	Filter   *sensitive.Filter

	UserModel   model.UsersModel
	BlacksModel model.BlacksModel

	Err error
}

func NewServiceContext(c config.Config) *ServiceContext {
	dbconn := sqlx.NewMysql(c.MysqlDb.DataSource)
	return &ServiceContext{
		Config:      c,
		Auth:        middleware.NewAuthMiddleware(c.Auth.AccessSecret).Handle,
		UserModel:   model.NewUsersModel(dbconn),
		BlacksModel: model.NewBlacksModel(dbconn),
	}
}

func (s *ServiceContext) InitRedis() {
	if s.Err != nil {
		return
	}
	s.RedisCli = utils.NewRedisClient(s.Config.Redis.Host, s.Config.Redis.Pass, 0)
	err := s.RedisCli.Ping(context.Background())
	if err != nil {
		logx.Error("failed to connect redis, err: ", err)
	}
}

func (s *ServiceContext) InitFilterFile() {
	if s.Err != nil {
		return
	}
	s.Filter = sensitive.New()
	err := s.Filter.LoadWordDict(s.Config.FilterFile)
	if err != nil {
		logx.Error("failed to load filter file, err: ", err)
	}
}
