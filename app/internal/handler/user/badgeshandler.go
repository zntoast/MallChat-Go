package user

import (
	"net/http"

	"mallchat-go/app/internal/pkg/common/result"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallchat-go/app/internal/logic/user"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"
)

func BadgesHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.BadgesReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewBadgesLogic(r.Context(), svcCtx)
		resp, err := l.Badges(&req)
		result.HttpResult(r, w, resp, err)
	}
}
