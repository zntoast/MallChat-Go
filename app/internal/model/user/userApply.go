package user

import "gorm.io/gorm"

// 用户申请表
type UserApply struct {
	gorm.Model
	Uid        int64  `gorm:"not null;comment:申请人UID"`
	Type       int    `gorm:"not null;comment:申请类型 1加好友"`
	TargetID   int64  `gorm:"not null;comment:接收人UID"`
	Msg        string `gorm:"not null;comment:申请信息"`
	Status     int    `gorm:"not null;comment:申请状态 1待审批 2同意"`
	ReadStatus int    `gorm:"not null;comment:阅读状态 1未读 2已读"`
}

const (
	UserApplyStatusWait  = 1 //  待审批
	UserApplyStatusAgree = 2 // 同意
)

const (
	UserApplyReadStatusUnread = 1 // 未读
	UserApplyReadStatusRead   = 2 // 已读
)
