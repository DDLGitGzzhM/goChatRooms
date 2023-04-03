package dao

import (
	"gorm.io/gorm"
	"test/models"
	"test/utils/database"
)

// CreateUser
// 用于用户的注册
func CreateUser(user models.User) *gorm.DB {
	return database.DB.Create(&user)
}
