package dao

import (
	"gorm.io/gorm"
	"test/models"
	"test/utils/database"
)

// CreateUserContact
// 创建用户之间的关系
func CreateUserConact(usercontact models.UserContact) *gorm.DB {
	return database.DB.Create(&usercontact)
}

// FindUserContactByOwnerIdAndTargetId
// 用于寻找 , 判断之间的关系
func FindUserContactByOwnerIdAndTargetId(OwnerId, TargetId int) models.UserContact {
	usercontact := models.UserContact{}
	database.DB.Where("Aid = ? and Bid = ?", OwnerId, TargetId).First(&usercontact)
	return usercontact
}

// ChangeFriendContact
// 更新用户之间的关系
func ChangeFrinedContact(usercontact models.UserContact) *gorm.DB {
	return database.DB.Model(&usercontact).Where("Aid = ? and Bid = ?", usercontact.OwnerId, usercontact.TargetId).Updates(
		models.UserContact{
			Type:  usercontact.Type,
			MTime: usercontact.MTime,
		})
}
