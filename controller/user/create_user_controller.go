package user

import (
	"fmt"
	"math/rand"
	"test/dao"
	"test/models"
	"test/utils"
	"time"
)

// 用于 返回 用户CRUD时候的错误
// 进行用户的格式规范
type UserError struct {
	//  错误码
	ErrorCode int
	//	错误信息
	ErrorMessage string
}

func (e *UserError) Error() string {
	return fmt.Sprintf("error code: %d, message: %s", e.ErrorCode, e.ErrorMessage)
}
func AddUser(name, password string) (err error) {
	temp, _ := FindUserByName(name)

	if temp == nil {
		return &UserError{
			ErrorCode:    1,
			ErrorMessage: "不能重复创建",
		}
	}
	if len(name) > 20 || name == "" {
		return &UserError{
			ErrorCode:    1,
			ErrorMessage: "用户名太长或者太短",
		}
	}
	if len(password) >= 100 || name == "" {
		return &UserError{
			ErrorCode:    1,
			ErrorMessage: "密码太长 或者 太短",
		}
	}
	user := models.User{}
	user.Name = name
	user.Salt = fmt.Sprintf("%06d", rand.Int31())
	user.Password = utils.MakePassword(password, user.Salt)
	user.Status = "上线"

	user.CTime = time.Now()
	user.MTime = time.Now()
	user.IsDelete = false

	dao.CreateUser(user)
	return &UserError{
		ErrorCode:    0,
		ErrorMessage: "创建成功",
	}
}
