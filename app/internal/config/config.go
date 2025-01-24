package config

import "github.com/zeromicro/go-zero/rest"

type Config struct {
	rest.RestConf
	MysqlDb struct {
		DataSource string
	}
	Redis struct {
		Host string
		Pass string
	}
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	FilterFile string
}
