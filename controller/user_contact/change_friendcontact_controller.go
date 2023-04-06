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
		CreateUserContact(TargetId, OwnerId, 1)
	}

	// 然后再进行更改
	//	对应的类型
	//  0表示还没创建, 但是其实是加白, 不过还没考虑如何将0进行加白
	//  1表示主动加白
	//  2表示拉黑
	//  3表示被拉黑
	var ownerType int
	var targetType int
	if Type == 1 {
		ownerType = 1
		targetType = 1
	} else {
		ownerType = 2
		targetType = 3
	}
	// 关系是双方的
	// 这里是自己给对方拉黑 ,
	usercontact := models.UserContact{
		OwnerId:  OwnerId,
		TargetId: TargetId,
		Type:     ownerType,
		MTime:    time.Now(),
	}
	dao.ChangeFrinedContact(usercontact)

	//这里是别人（相当于现在的自己)
	//和自己是被拉黑的关系
	usercontact = models.UserContact{
		OwnerId:  TargetId,
		TargetId: OwnerId,
		Type:     targetType,
		MTime:    time.Now(),
	}
	dao.ChangeFrinedContact(usercontact)

	return UserContactError{
		ErrorCode:    0,
		ErrorMessage: "拉黑成功",
	}
}
