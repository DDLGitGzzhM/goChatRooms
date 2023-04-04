package user_contact

import (
	"fmt"
	"test/dao"
)

// 主要用于控制
// 用户关系创建的判断
func FindUserContactByOwnerIdAndTargetId(OwnerId, TargetId int) (Type int, err error) {
	usercontact := dao.FindUserContactByOwnerIdAndTargetId(OwnerId, TargetId)
	fmt.Println("usercontact >>>>>>>", usercontact)
	return usercontact.Type, UserContactError{
		ErrorCode:    0,
		ErrorMessage: "查询成功",
	}
}
