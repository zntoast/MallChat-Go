package ws

import (
	"net/http"

	"mallchat-go/app/internal/logic/ws"
	"mallchat-go/app/internal/svc"
)

func ConnectHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	l := ws.NewConnectLogic(svcCtx)
	return l.Connect
}
