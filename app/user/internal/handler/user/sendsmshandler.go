package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallchat-go/app/user/internal/logic/user"
	"mallchat-go/app/user/internal/svc"
	"mallchat-go/app/user/internal/types"
)

// 发送验证码
func SendSmsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.SendSmsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewSendSmsLogic(r.Context(), svcCtx)
		resp, err := l.SendSms(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
