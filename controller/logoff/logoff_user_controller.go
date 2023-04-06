package logoff

import (
	"test/controller/user"
	"test/dao"
	"test/models"
	"time"
)

func Logoff(name string) error {
	user_t := models.User{}
	temp, _ := user.FindUserByName(name)
	user_t.Status = "下线"
	user_t.ID = temp.ID
	user_t.IsDelete = true
	user_t.MTime = time.Now()
	dao.UpdateUser(user_t)

	return &user.UserError{
		ErrorCode:    0,
		ErrorMessage: "注销成功",
	}
}
