// Code generated by goctl. DO NOT EDIT.
// goctl 1.7.2

package handler

import (
	"net/http"

	login "mallchat-go/app/internal/handler/login"
	message "mallchat-go/app/internal/handler/message"
	user "mallchat-go/app/internal/handler/user"
	"mallchat-go/app/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				// login
				Method:  http.MethodPost,
				Path:    "/login",
				Handler: login.LoginHandler(serverCtx),
			},
			{
				// register
				Method:  http.MethodPost,
				Path:    "/register",
				Handler: login.RegisterHandler(serverCtx),
			},
		},
		rest.WithPrefix("/capi/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					// 获取消息列表
					Method:  http.MethodPost,
					Path:    "/message/list",
					Handler: message.GetMessageListHandler(serverCtx),
				},
				{
					// 发送消息
					Method:  http.MethodPost,
					Path:    "/message/send",
					Handler: message.SendMessageHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/capi/im"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				// 重置密码
				Method:  http.MethodPost,
				Path:    "/password/reset",
				Handler: user.ResetPasswordHandler(serverCtx),
			},
			{
				// 发送验证码
				Method:  http.MethodPost,
				Path:    "/sms/send",
				Handler: user.SendSmsHandler(serverCtx),
			},
		},
		rest.WithPrefix("/capi/user"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.Auth},
			[]rest.Route{
				{
					// 上传头像
					Method:  http.MethodPost,
					Path:    "/avatar/upload",
					Handler: user.UploadAvatarHandler(serverCtx),
				},
				{
					// 获取用户信息
					Method:  http.MethodGet,
					Path:    "/info",
					Handler: user.GetUserInfoHandler(serverCtx),
				},
				{
					// 更新用户信息
					Method:  http.MethodPost,
					Path:    "/update",
					Handler: user.UpdateUserHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/capi/user"),
	)
}