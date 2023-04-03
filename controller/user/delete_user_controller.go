package user

import "test/dao"

func DeleteUser(name string) error {
	user, _ := FindUserByName(name)
	//判断是否能删
	if user.Name == "" {
		return &UserError{
			ErrorCode:    1,
			ErrorMessage: "删除失败 ,因为用户不存在",
		}
	}
	dao.DeleteUser(user)
	return &UserError{
		ErrorCode:    0,
		ErrorMessage: "删除成功",
	}
}
