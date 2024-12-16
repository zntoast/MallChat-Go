package middleware

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"mallchat-go/app/user/internal/utils"

	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
)

var ErrUnauthorized = errors.New("未授权")

func NewAuth(secret string) rest.Middleware {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			auth := r.Header.Get("Authorization")
			if auth == "" {
				httpx.Error(w, ErrUnauthorized)
				return
			}

			parts := strings.SplitN(auth, " ", 2)
			if !(len(parts) == 2 && parts[0] == "Bearer") {
				httpx.Error(w, ErrUnauthorized)
				return
			}

			claims, err := utils.ParseToken(parts[1], secret)
			if err != nil {
				httpx.Error(w, ErrUnauthorized)
				return
			}

			// 将用户ID添加到请求头中
			r.Header.Set("X-User-ID", strconv.FormatInt(claims.UserId, 10))
			next(w, r)
		}
	}
}
