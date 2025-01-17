package user

import "gorm.io/gorm"

// 用户背包表
type UserBackpack struct {
	gorm.Model
	Uid        int64  `gorm:"not null;comment:用户UID"`
	ItemID     int64  `gorm:"not null;comment:物品ID"`
	Status     int    `gorm:"not null;comment:使用状态 0.未失效 1失效"`
	Idempotent string `gorm:"not null;comment:幂等号"`
}

const (
	UserBackpackStatusInvalid = 1 // 失效
	UserBackpackStatusValid   = 0 // 未失效
)
