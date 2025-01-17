package user

import (
	"time"

	"gorm.io/gorm"
)

// 用户表
type User struct {
	gorm.Model
	Name         string    `gorm:"not null;comment:用户昵称"`
	Avatar       string    `gorm:"not null;comment:用户头像"`
	Sex          int       `gorm:"not null;comment:性别 1为男性，2为女性"`
	OpenID       string    `gorm:"not null;comment:微信openid用户标识"`
	ActiveStatus int       `gorm:"not null;comment:上下线状态 1在线 2离线"`
	LastOptTime  time.Time `gorm:"not null;comment:最后上下线时间"`
	ItemID       int64     `gorm:"not null;comment:佩戴的徽章ID"`
	Status       int       `gorm:"not null;comment:用户状态 0正常 1拉黑"`
}

const (
	UserActiveStatusOnline  = 1 // 在线
	UserActiveStatusOffline = 2 // 离线
)

const (
	UserSexMale   = 1 // 男性
	UserSexFemale = 2 // 女性
)

const (
	UserStatusNormal  = 0 // 正常
	UserStatusBlacked = 1 // 拉黑
)
