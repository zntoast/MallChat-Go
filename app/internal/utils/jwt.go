package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type Claims struct {
	UserId int64 `json:"userId"`
	jwt.RegisteredClaims
}

type JWTUtils struct {
	secretKey []byte
}

func NewJWTUtils(secretKey string) *JWTUtils {
	return &JWTUtils{
		secretKey: []byte(secretKey),
	}
}

func (j *JWTUtils) GenerateToken(userId int64) (string, int64, int64, error) {
	now := time.Now()
	expireTime := now.Add(2 * time.Hour)
	refreshTime := now.Add(1 * time.Hour)

	claims := Claims{
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expireTime),
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.secretKey)
	if err != nil {
		return "", 0, 0, err
	}

	return tokenString, expireTime.Unix(), refreshTime.Unix(), nil
}

func (j *JWTUtils) ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return j.secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, jwt.ErrSignatureInvalid
}
