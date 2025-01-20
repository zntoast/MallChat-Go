package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallchat-go/app/internal/logic/user"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"
)

// 修改用户名
func ModifyNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyNameReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewModifyNameLogic(r.Context(), svcCtx)
		err := l.ModifyName(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
