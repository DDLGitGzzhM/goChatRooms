package login

import (
	user "test/controller/user"
	"test/dao"
	"test/models"
	"test/utils"
	"time"
)

func Login(name, password string) error {
	temp, _ := user.FindUserByName(name)

	if temp.Name == "" {
		return &user.UserError{
			ErrorCode:    1,
			ErrorMessage: "用户不存在",
		}
	}

	if temp.IsDelete == true {
		return &user.UserError{
			ErrorCode:    1,
			ErrorMessage: "用户已注销",
		}
	}

	if utils.ValidPassword(password, temp.Salt, temp.Password) == false {
		return &user.UserError{
			ErrorCode:    1,
			ErrorMessage: "密码错误",
		}
	}

	//将下线信息修改成上线
	user_t := models.User{}
	temp, _ = user.FindUserByName(name)
	user_t.ID = temp.ID
	user_t.Status = "上线"
	user_t.IsDelete = false
	user_t.MTime = time.Now()
	dao.UpdateUser(user_t)

	return &user.UserError{
		ErrorCode:    0,
		ErrorMessage: "登录成功",
	}
}
