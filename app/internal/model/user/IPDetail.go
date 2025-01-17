package user

import "gorm.io/gorm"

// 用户IP信息表
type IPDetail struct {
	gorm.Model
	IP        string `gorm:"not null;comment:注册时的ip"`
	ISP       string `gorm:"not null;comment:最新登录的ip"`
	ISPID     string `gorm:"not null;comment:ISP ID"`
	City      string `gorm:"not null;comment:城市"`
	CityID    string `gorm:"not null;comment:城市ID"`
	Country   string `gorm:"not null;comment:国家"`
	CountryID string `gorm:"not null;comment:国家ID"`
	Region    string `gorm:"not null;comment:地区"`
	RegionID  string `gorm:"not null;comment:地区ID"`
	Ipid      int64  `gorm:"not null;comment:用户IPID"`
}
