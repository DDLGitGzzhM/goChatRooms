package login

import (
	user "test/controller/user"
	"test/utils"
)

func Login(name, password string) error {
	temp, _ := user.FindUserByName(name)

	if temp.Name == "" {
		return &user.UserError{
			ErrorCode:    1,
			ErrorMessage: "用户不存在",
		}
	}

	if utils.ValidPassword(password, temp.Salt, temp.Password) == false {
		return &user.UserError{
			ErrorCode:    1,
			ErrorMessage: "密码错误",
		}
	}

	return &user.UserError{
		ErrorCode:    0,
		ErrorMessage: "登录成功",
	}
}
