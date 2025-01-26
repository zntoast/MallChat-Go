package result

import (
	"net/http"

	"mallchat-go/app/internal/errors"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
)

const (
	BaseCode = 100000
)

func mapStatusCode(code uint32) int {
	return int(code) / BaseCode
}

// http返回
func HttpResult(r *http.Request, w http.ResponseWriter, resp interface{}, err error) {
	if err == nil {
		r := Success(resp)
		httpx.WriteJson(w, http.StatusOK, r)
	} else {
		WriteError(w, r, err)
	}
}

func WriteError(w http.ResponseWriter, r *http.Request, e error) {
	var errcode errors.ErrorCode
	var errmsg string
	if err, ok := e.(*errors.ErrorEntry); ok { //自定义错误类型
		errcode = err.GetErrCode()
		errmsg = err.GetErrMsg()
	} else {
		e = errors.Adapt(e)
		errmsg = "服务器开了点小差，请稍后再试"
	}
	logx.WithContext(r.Context()).Errorf("【API-ERR】 : %+v ", e)
	httpStatus := mapStatusCode(uint32(errcode))
	httpx.WriteJson(w, httpStatus, Error(uint32(errcode), errmsg))
}

// http 参数错误返回
func ParamErrorResult(r *http.Request, w http.ResponseWriter, err error) {
	httpx.WriteJson(w, mapStatusCode(uint32(errors.ErrStatusBadRequest)), Error(uint32(errors.ErrStatusBadRequest), err.Error()))
}
