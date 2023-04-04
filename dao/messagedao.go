package dao

import (
	"gorm.io/gorm"
	"test/models"
	"test/utils/database"
)

// 那么这里采用 发一次就存一次的形式咯
func AddMessage(message models.Message) *gorm.DB {
	return database.DB.Create(&message)
}

func GetMessageListAny() []*models.Message {
	message := make([]*models.Message, 0, 50)
	database.DB.Find(&message)
	return message
}

// 首先 问题是 toUserIds是一个json
// 懂了 toUserIds 其实是一个冗余的 变量
// 所以 我们需要引入 TargetId
func GetMessageListNotAny(OwnerId, TargetId int) []*models.Message {
	message := make([]*models.Message, 0, 50)
	database.DB.Where("userId = ? and targetId = ?", OwnerId, TargetId).Find(&message)
	return message
}
