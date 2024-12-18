package svc

import (
	"mallchat-go/app/internal/config"
	"mallchat-go/app/internal/middleware"
	"mallchat-go/app/internal/model"
	wsTypes "mallchat-go/app/internal/types/ws"
	"mallchat-go/app/internal/utils"
	wsImpl "mallchat-go/app/internal/ws"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/rest"
)

type ServiceContext struct {
	Config              config.Config
	Auth                rest.Middleware
	JWT                 *utils.JWTUtils
	UserModel           model.UserModel
	MessageModel        model.MessageModel
	GroupModel          model.GroupModel
	OfflineMessageModel model.OfflineMessageModel
	Redis               *utils.RedisClient
	OSS                 *utils.OSSClient
	WS                  wsTypes.Manager
	Models              *model.Models
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.MySQL.DataSource)
	jwt := utils.NewJWTUtils(c.Auth.AccessSecret)
	redis := utils.NewRedisClient(c.Redis.Host, c.Redis.Password, c.Redis.DB)
	oss, err := utils.NewOSSClient(
		c.OSS.Endpoint,
		c.OSS.AccessKey,
		c.OSS.AccessSecret,
		c.OSS.Bucket,
		c.OSS.BucketDomain,
	)
	if err != nil {
		panic(err)
	}

	cacheConf := c.Cache
	models := model.NewModels(
		model.NewGroupModel(conn, cacheConf),
		model.NewOfflineMessageModel(conn, cacheConf),
	)

	ctx := &ServiceContext{
		Config:    c,
		UserModel: model.NewUserModel(conn, cacheConf),

		MessageModel:        model.NewMessageModel(conn, cacheConf),
		GroupModel:          model.NewGroupModel(conn, cacheConf),
		OfflineMessageModel: model.NewOfflineMessageModel(conn, cacheConf),
		Redis:               redis,
		OSS:                 oss,
		Models:              models,
	}

	ctx.WS = wsImpl.NewManager(models)
	ctx.Auth = middleware.NewAuthMiddleware(jwt).Handle
	return ctx
}
