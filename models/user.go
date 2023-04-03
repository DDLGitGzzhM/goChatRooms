package models

import "time"

type User struct {
	ID       int
	Name     string
	Password string
	Status   string    `gorm:"column:status"`
	IsDelete bool      `gorm:"column:isDelete"`
	CTime    time.Time `gorm:"column:cTime"`
	MTime    time.Time `gorm:"column:mTime"`
}

func (User) TableName() string {
	return "user"
}
