package middleware

import (
	"context"
	"net/http"
	"strings"

	"mallchat-go/app/internal/common"
	"mallchat-go/app/internal/utils"
)

type AuthMiddleware struct {
	jwt *utils.JWTUtils
}

func NewAuthMiddleware(jwt *utils.JWTUtils) *AuthMiddleware {
	return &AuthMiddleware{
		jwt: jwt,
	}
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "未授权", http.StatusUnauthorized)
			return
		}

		parts := strings.SplitN(authHeader, " ", 2)
		if !(len(parts) == 2 && parts[0] == "Bearer") {
			http.Error(w, "无效的认证格式", http.StatusUnauthorized)
			return
		}

		token := parts[1]
		claims, err := m.jwt.ParseToken(token)
		if err != nil {
			http.Error(w, "无效的token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), common.UserIDKey, claims.UserId)
		r = r.WithContext(ctx)

		next(w, r)
	}
}
