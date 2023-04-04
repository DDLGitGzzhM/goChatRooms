package user

import (
	"test/dao"
)

func GetUserList() ([]string, error) {
	users := dao.GetUserList("上线")

	// 这里传一个切片
	// sname 切片用于存储上线用户
	// 因为不会proto 里面存储user 所以使用[]string
	sname := make([]string, 0)
	for _, val := range users {
		sname = append(sname, val.Name)
	}
	return sname, &UserError{
		ErrorCode:    0,
		ErrorMessage: "查询用户列表成功",
	}
}
