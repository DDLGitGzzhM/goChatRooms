package user_contact

import (
	"fmt"
	"test/dao"
	"test/models"
	"time"
)

func ChangeFriendContact(OwnerId, TargetId, Type int) error {
	temp, _ := FindUserContactByOwnerIdAndTargetId(OwnerId, TargetId)
	fmt.Println("temp >>>>>>>", temp)
	// 如果并没有关系 , 我们需要先创建关系
	// 不过这个会有影响 , 也就是第一次创建的时间和 "主动" 修改的时间一样
	// 可是这个创建时间的作用是什么 ？
	if temp == 0 {
		CreateUserContact(OwnerId, TargetId, 1)
	}

	// 然后再进行更改
	usercontact := models.UserContact{
		OwnerId:  OwnerId,
		TargetId: TargetId,
		Type:     Type,
		MTime:    time.Now(),
	}

	dao.ChangeFrinedContact(usercontact)

	return UserContactError{
		ErrorCode:    0,
		ErrorMessage: "更改成功",
	}
}
