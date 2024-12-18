package utils

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisClient struct {
	client *redis.Client
}

func NewRedisClient(host, password string, db int) *RedisClient {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})

	return &RedisClient{
		client: client,
	}
}

// SetVerifyCode 设置验证码
func (r *RedisClient) SetVerifyCode(ctx context.Context, mobile, scene, code string) error {
	key := fmt.Sprintf("verify_code:%s:%s", scene, mobile)
	return r.client.Set(ctx, key, code, 5*time.Minute).Err()
}

// GetVerifyCode 获取验证码
func (r *RedisClient) GetVerifyCode(ctx context.Context, mobile, scene string) (string, error) {
	key := fmt.Sprintf("verify_code:%s:%s", scene, mobile)
	return r.client.Get(ctx, key).Result()
}

// DelVerifyCode 删除验证码
func (r *RedisClient) DelVerifyCode(ctx context.Context, mobile, scene string) error {
	key := fmt.Sprintf("verify_code:%s:%s", scene, mobile)
	return r.client.Del(ctx, key).Err()
}

// SetUserToken 设置用户token
func (r *RedisClient) SetUserToken(ctx context.Context, userId int64, token string) error {
	key := fmt.Sprintf("user_token:%d", userId)
	return r.client.Set(ctx, key, token, 24*time.Hour).Err()
}

// GetUserToken 获取用户token
func (r *RedisClient) GetUserToken(ctx context.Context, userId int64) (string, error) {
	key := fmt.Sprintf("user_token:%d", userId)
	return r.client.Get(ctx, key).Result()
}

// DelUserToken 删除用户token
func (r *RedisClient) DelUserToken(ctx context.Context, userId int64) error {
	key := fmt.Sprintf("user_token:%d", userId)
	return r.client.Del(ctx, key).Err()
}
