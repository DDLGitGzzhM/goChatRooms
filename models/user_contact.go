package models

import "time"

type UserContact struct {
	//  用户自己的 id
	OwnerId int `gorm:"column:Aid"`
	//	对应的ID
	TargetId int `gorm:"column:Bid"`
	//	对应的类型
	//  0表示还没创建, 但是其实是加白, 不过还没考虑如何将0进行加白
	//  1表示主动加白
	//  2表示拉黑
	//  3表示被拉黑
	Type int `gorm:"column:contact_Type"`

	CTime time.Time `gorm:"column:cTime"`
	MTime time.Time `gorm:"column:mTime"`
}

func (table *UserContact) TableName() string {
	return "user_contact"
}
