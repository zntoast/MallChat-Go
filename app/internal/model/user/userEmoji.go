package user

import "gorm.io/gorm"

// 用户表情包表
type UserEmoji struct {
	gorm.Model
	Uid           int64  `gorm:"not null;comment:用户UID"`
	ExpressionURL string `gorm:"not null;comment:表情地址"`
	DeleteStatus  int    `gorm:"not null;default:0;comment:逻辑删除(0-正常,1-删除)"`
}

const (
	UserEmojiDeleteStatusNormal  = 0
	UserEmojiDeleteStatusDeleted = 1
)
