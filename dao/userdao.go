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

// GetUserList
// 群聊用于获取所有人数 并且 判断是否在线
func GetUserList(status string) (users []models.User) {
	database.DB.Where("status = ?", status).Find(&users)
	return
}

// FindUserByName
// 用于检测注册是否有重复的用户
func FindUserByName(name string) models.User {
	user := models.User{}
	database.DB.Where("name = ?", name).First(&user)
	return user
}

// DeleteUser
// 删除用户真实的删除
func DeleteUser(user models.User) *gorm.DB {
	return database.DB.Delete(&user)
}

// 更新用户
// 这里可以拿来做逻辑上的删除 , 以及更新更新时间
// 更新上下线状态
func UpdateUser(user models.User) *gorm.DB {
	return database.DB.Model(&user).Updates(
		models.User{
			ID:       user.ID,
			Status:   user.Status,
			IsDelete: user.IsDelete,
			MTime:    user.MTime,
		})
}
