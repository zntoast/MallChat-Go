package login

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallchat-go/app/user/internal/logic/login"
	"mallchat-go/app/user/internal/svc"
	"mallchat-go/app/user/internal/types"
)

// login
func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UserLoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := login.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
