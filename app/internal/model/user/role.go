package user

import "gorm.io/gorm"

// 角色表
type Role struct {
	gorm.Model
	Name string `gorm:"not null;comment:角色名称"`
}

// 用户角色关系表
type UserRole struct {
	gorm.Model
	Uid    int64 `gorm:"not null;comment:用户UID"`
	RoleID int64 `gorm:"not null;comment:角色ID"`
}
