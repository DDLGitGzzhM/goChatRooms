package user

import (
	"test/dao"
	"test/models"
)

func FindUserByName(name string) (user *models.User, err error) {
	temp := dao.FindUserByName(name)
	user = &temp
	return user, &UserError{
		ErrorCode:    0,
		ErrorMessage: "查询成功",
	}
}
