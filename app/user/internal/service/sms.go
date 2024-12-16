package service

import (
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/stores/redis"
)

type SmsService struct {
	rds *redis.Redis
}

func NewSmsService(rds *redis.Redis) *SmsService {
	return &SmsService{rds: rds}
}

func (s *SmsService) GetVerifyCode(ctx context.Context, mobile, scene string) (string, error) {
	key := fmt.Sprintf("sms:%s:%s", scene, mobile)
	return s.rds.Get(key)
}

func (s *SmsService) SetVerifyCode(ctx context.Context, mobile, scene, code string) error {
	key := fmt.Sprintf("sms:%s:%s", scene, mobile)
	return s.rds.Setex(key, code, int(5*time.Minute.Seconds())) // 5分钟有效期
}

func (s *SmsService) CheckSendFrequency(ctx context.Context, mobile string) error {
	key := fmt.Sprintf("sms:frequency:%s", mobile)
	exists, err := s.rds.Exists(key)
	if err != nil {
		return err
	}
	if exists {
		return fmt.Errorf("发送太频繁，请稍后再试")
	}

	// 设置60秒内不能重复发送
	return s.rds.Setex(key, "1", 60)
}
