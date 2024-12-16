package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	DB struct {
		DataSource string
	}
	Cache cache.CacheConf
	Auth  struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Host string
		Pass string
	}
	Upload struct {
		SaveDir string
		MaxSize int64  // 单位MB
		BaseUrl string // 访问URL前缀
	}
}
