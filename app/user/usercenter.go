package main

import (
	"flag"
	"fmt"
	"net/http"
	"strings"

	"mallchat-go/app/user/internal/config"
	"mallchat-go/app/user/internal/handler"
	"mallchat-go/app/user/internal/middleware"
	"mallchat-go/app/user/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/usercenter.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	// 添加认证中间件
	authMiddleware := middleware.NewAuth(c.Auth.AccessSecret)
	server.Use(func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			// 只对需要认证的路径应用中间件
			if strings.HasPrefix(r.URL.Path, "/capi/user/info") ||
				strings.HasPrefix(r.URL.Path, "/capi/user/update") ||
				strings.HasPrefix(r.URL.Path, "/capi/user/avatar") {
				authMiddleware(next)(w, r)
				return
			}
			next(w, r)
		}
	})

	// 添加静态文件服务
	fs := http.FileServer(http.Dir(c.Upload.SaveDir))
	server.AddRoute(rest.Route{
		Method:  http.MethodGet,
		Path:    "/static/{filepath:.+}",
		Handler: http.StripPrefix("/static/", fs).ServeHTTP,
	})

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 注册中间件
	server.Use(middleware.ErrorHandler)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
