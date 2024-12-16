package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JwtClaims struct {
	UserId int64 `json:"userId"`
	jwt.RegisteredClaims
}

func GenerateToken(secretKey string, userId int64, expireDuration time.Duration) (string, int64, error) {
	expireAt := time.Now().Add(expireDuration)
	claims := JwtClaims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(secretKey))
	return tokenString, expireAt.Unix(), err
}
