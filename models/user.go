package models

import "time"

type User struct {
	//基本字段
	ID       int
	Name     string
	Password string
	Status   string    `gorm:"column:status"`
	IsDelete bool      `gorm:"column:isDelete"`
	CTime    time.Time `gorm:"column:cTime"`
	MTime    time.Time `gorm:"column:mTime"`
	//用于加密
	Salt string
}

func (User) TableName() string {
	return "user"
}
