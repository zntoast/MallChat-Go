package middleware

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func ErrorHandler(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		next(w, r)
	}
}

func ResponseHandler(w http.ResponseWriter, r *http.Request, resp interface{}, err error) {
	if err != nil {
		httpx.WriteJson(w, http.StatusOK, Response{
			Code: -1,
			Msg:  err.Error(),
		})
		return
	}

	httpx.WriteJson(w, http.StatusOK, Response{
		Code: 0,
		Msg:  "success",
		Data: resp,
	})
}
