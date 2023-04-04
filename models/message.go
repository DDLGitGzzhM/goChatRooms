package models

import "time"

type Message struct {
	//  基本字段
	ID        int
	UserId    int       `gorm:"column:userId"`
	RoomId    int       `gorm:"column:roomId"`
	IsSendAll bool      `gorm:"column:isSendAll"`
	ToUserIds string    `gorm:"column:toUserIds"`
	Content   string    `gorm:"column:content"`
	CTime     time.Time `gorm:"column:cTime"`
	MTime     time.Time `gorm:"column:mTime"`
	//
}

func (Message) TableName() string {
	return "message"
}
