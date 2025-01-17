package user

import "gorm.io/gorm"

// 功能物品配置表
type ItemConfig struct {
	gorm.Model
	Type     int    `gorm:"not null;comment:物品类型 1改名卡 2徽章"`
	Img      string `gorm:"not null;comment:物品图片"`
	Describe string `gorm:"not null;comment:物品功能描述"`
}
