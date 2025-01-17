package user

import "gorm.io/gorm"

// 用户IP信息表
type UserIPInfo struct {
	gorm.Model
	CreateIP       string `gorm:"not null;comment:注册时的ip"`
	CreateIPDetail int64  `gorm:"not null;comment:注册时的ip详情ID"`
	UpdateIP       string `gorm:"not null;comment:最新登录的ip"`
	UpdateIPDetail int64  `gorm:"not null;comment:最新登录的ip详情ID"`
	Uid            int64  `gorm:"not null;comment:用户UID"`
}
