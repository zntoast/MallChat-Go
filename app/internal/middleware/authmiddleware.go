package middleware

import (
	"context"
	"errors"
	"fmt"
	"mallchat-go/app/internal/model"
	"mallchat-go/app/internal/pkg/utils"
	"mallchat-go/app/internal/svc"
	"net/http"
	"strings"
)

type UserProvider interface {
	GetUserById(ctx context.Context, userId int64) (*model.Users, error)
}

type AuthMiddleware struct {
	secret string
	svcCtx *svc.ServiceContext
}

func NewAuthMiddleware(svcCtx *svc.ServiceContext) *AuthMiddleware {
	return &AuthMiddleware{
		secret: svcCtx.Config.Auth.AccessSecret,
		svcCtx: svcCtx,
	}
}

type AuthServiceResp struct{}

func GetAuthRespFromCtx(ctx context.Context) model.Users {
	return ctx.Value(AuthServiceResp{}).(model.Users)
}

func (m *AuthMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Token validation
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			writeError(w, http.StatusUnauthorized, errors.New("authorization header is required"))
			return
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			writeError(w, http.StatusUnauthorized, errors.New("invalid authorization header format"))
			return
		}
		token := tokenParts[1]

		claims, err := utils.ParseToken(token, m.secret)
		if err != nil {
			writeError(w, http.StatusUnauthorized, fmt.Errorf("invalid token: %v", err))
			return
		}

		userPtr, err := m.svcCtx.UserModel.GetUserById(r.Context(), uint64(claims.UserId))
		if err != nil || userPtr == nil {
			writeError(w, http.StatusUnauthorized, errors.New("user not found"))
			return
		}

		r2 := r.WithContext(context.WithValue(r.Context(), AuthServiceResp{}, *userPtr))
		next(w, r2)
	}
}

func writeError(w http.ResponseWriter, code int, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write([]byte(fmt.Sprintf(`{"error": "%s"}`, err.Error())))
}
