package message

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mallchat-go/app/internal/logic/message"
	"mallchat-go/app/internal/svc"
	"mallchat-go/app/internal/types"
)

// 获取消息列表
func GetMessageListHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GetMessageListReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := message.NewGetMessageListLogic(r.Context(), svcCtx)
		resp, err := l.GetMessageList(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
