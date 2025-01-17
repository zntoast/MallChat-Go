package user

import "gorm.io/gorm"

// 黑名单表
type Black struct {
	gorm.Model
	Type   int    `gorm:"not null;comment:拉黑目标类型 1.ip 2.uid"`
	Target string `gorm:"not null;comment:拉黑目标"`
}
