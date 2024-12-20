package config

import (
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
	Auth struct {
		AccessSecret string
	}
	MySQL struct {
		DataSource string
	}
	Redis struct {
		Host     string
		Password string
		DB       int
	}
	OSS struct {
		Endpoint     string
		AccessKey    string
		AccessSecret string
		Bucket       string
		BucketDomain string // 用于生成访问URL
	}
	CacheRedis cache.ClusterConf
}
