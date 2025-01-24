package main

import (
	"flag"
	"fmt"

	"mallchat-go/app/internal/config"
	"mallchat-go/app/internal/handler"
	"mallchat-go/app/internal/middleware"
	"mallchat-go/app/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/mallchatgo.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	ctx.InitRedis()      // 初始化redis连接
	ctx.InitFilterFile() // 初始化过滤词库
	ctx.Auth = middleware.NewAuthMiddleware(ctx).Handle
	if ctx.Err != nil {
		logx.ErrorStack("Starting server error ", ctx.Err)
		return
	}

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
