package svc

import (
	"fmt"
	"mallchat-go/app/internal/config"
	"mallchat-go/app/internal/middleware"
	modelUser "mallchat-go/app/internal/model/user"

	"github.com/zeromicro/go-zero/rest"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type ServiceContext struct {
	Config config.Config
	Auth   rest.Middleware
	Db     *gorm.DB
	Err    error
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Auth:   middleware.NewAuthMiddleware(c.Auth.AccessSecret).Handle,
	}
}

func (s *ServiceContext) IninMysqlDB() {
	if s.Err != nil {
		return
	}
	db, err := gorm.Open(mysql.Open(s.Config.MysqlDb.DataSource), &gorm.Config{
		DisableAutomaticPing: s.Config.MysqlDb.AutoPing,
		Logger:               logger.Default,
	})
	if err != nil {
		s.Err = fmt.Errorf("failed to connect database, err: %v", err)
		return
	}
	err = db.AutoMigrate(
		modelUser.Black{},
		modelUser.ItemConfig{},
		modelUser.Role{},
		modelUser.User{},
		modelUser.UserApply{},
		modelUser.UserBackpack{},
		modelUser.UserEmoji{},
		modelUser.UserFriend{},
		modelUser.UserRole{},
		modelUser.UserIPInfo{},
		modelUser.IPDetail{},
	)
	if err != nil {
		s.Err = fmt.Errorf("failed to migrate database, err: %v", err)
		return
	}
	s.Db = db
}
