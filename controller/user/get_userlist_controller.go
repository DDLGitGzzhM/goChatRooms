package user

import (
	"test/dao"
	"test/models"
)

func GetUserList() ([]models.User, error) {
	users := dao.GetUserList("上线")
	return users, &UserError{
		ErrorCode:    0,
		ErrorMessage: "查询用户列表成功",
	}
}
