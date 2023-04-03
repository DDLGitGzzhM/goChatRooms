package user

import (
	"test/dao"
)

func GetUserList() ([]string, error) {
	users := dao.GetUserList("上线")
	// 这里传一个切片 不过传的是名字
	sname := make([]string, 10)
	for _, val := range users {
		sname = append(sname, val.Name)
	}
	return sname, &UserError{
		ErrorCode:    0,
		ErrorMessage: "查询用户列表成功",
	}
}
