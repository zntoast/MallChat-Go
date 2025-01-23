package user

import (
	"net/http"

	"mallchat-go/app/internal/pkg/common/result"

	"mallchat-go/app/internal/logic/user"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func ModifyNameHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ModifyNameReq
		if err := httpx.Parse(r, &req); err != nil {
			result.ParamErrorResult(r, w, err)
			return
		}

		l := user.NewModifyNameLogic(r.Context(), svcCtx)
		err := l.ModifyName(&req)
		result.HttpResult(r, w, nil, err)
	}
}
