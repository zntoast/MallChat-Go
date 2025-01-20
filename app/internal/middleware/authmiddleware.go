package middleware

import (
	"context"
	"mallchat-go/app/internal/utils"
	"net/http"
	"strings"
)

type AuthMiddleware struct {
	secret string
}

func NewAuthMiddleware(secret string) *AuthMiddleware {
	return &AuthMiddleware{secret: secret}
}

type AuthServiceResp struct{}

func GetAuthRespFromCtx(ctx context.Context) int64 {
	return ctx.Value(AuthServiceResp{}).(int64)
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "auth header is missing", http.StatusUnauthorized)
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			http.Error(w, "invalid auth header", http.StatusUnauthorized)
			return
		}
		token := tokenParts[1]
		claims, err := utils.ParseToken(token, m.secret)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}
		r2 := r.WithContext(context.WithValue(r.Context(), AuthServiceResp{}, claims.UserId))

		next(w, r2)
	}
}
