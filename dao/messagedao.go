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
