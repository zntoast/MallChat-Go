package user

import (
	"net/http"

	"mallchat-go/app/internal/pkg/common/result"

	"mallchat-go/app/internal/logic/user"
	"mallchat-go/app/internal/svc"
)

func WearingBadgeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewWearingBadgeLogic(r.Context(), svcCtx)
		err := l.WearingBadge()
		result.HttpResult(r, w, nil, err)
	}
}
