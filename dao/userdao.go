package dao

import (
	"gorm.io/gorm"
	"test/models"
	"test/utils"
)

// CreateUser
// 用于用户的注册
func CreateUser(user models.User) *gorm.DB {
	return utils.DB.Create(&user)
}
