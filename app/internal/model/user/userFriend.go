package user

import "gorm.io/gorm"

// 用户联系人表
type UserFriend struct {
	gorm.Model
	Uid          int64 `gorm:"not null;comment:用户UID"`
	FriendUID    int64 `gorm:"not null;comment:好友UID"`
	DeleteStatus int   `gorm:"not null;default:0;comment:逻辑删除(0-正常,1-删除)"`
}

const (
	UserFriendDeleteStatusNormal  = 0
	UserFriendDeleteStatusDeleted = 1
)
