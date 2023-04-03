package logout

import (
	"test/controller/user"
	"test/dao"
	"test/models"
	"time"
)

func Logout(name string) error {
	user_t := models.User{}
	temp, _ := user.FindUserByName(name)

	user_t.ID = temp.ID
	user_t.Status = "下线"
	user_t.IsDelete = false
	user_t.MTime = time.Now()
	dao.UpdateUser(user_t)
	return &user.UserError{
		ErrorCode:    0,
		ErrorMessage: "下线成功",
	}
}
