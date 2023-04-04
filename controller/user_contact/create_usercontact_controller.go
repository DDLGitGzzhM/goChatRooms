package user_contact

import (
	"fmt"
	"test/dao"
	"test/models"
	"time"
)

// 用于 返回 用户CRUD时候的错误
// 进行用户的格式规范
type UserContactError struct {
	//  错误码
	ErrorCode int
	//	错误信息
	ErrorMessage string
}

// 可能User类这里是 *的关系, 导致有一个较为严重的错误
// 所以这里并不选择使用 *
func (e UserContactError) Error() string {
	return fmt.Sprintf("error code: %d, message: %s", e.ErrorCode, e.ErrorMessage)
}

// CreateUserContact
// 用于主动加白两者的关系

// 因为在Usercontact中加入了 Ctime字段
// 所以才会有此函数

// 认为这个函数可以放在ChangeContact之中进行执行
func CreateUserContact(OwnerId, TargetId, Type int) error {

	// 非法判断
	//if tempType > 0 {
	//	return UserContactError{
	//		ErrorCode:    1,
	//		ErrorMessage: "已经具有关系 , 不需要重复创建",
	//	}
	//}

	// 进行创建
	usercontact := models.UserContact{
		OwnerId:  OwnerId,
		TargetId: TargetId,
		Type:     Type,
		CTime:    time.Now(),
		MTime:    time.Now(),
	}
	dao.CreateUserConact(usercontact)

	return UserContactError{
		ErrorCode:    0,
		ErrorMessage: "已经成功创建两者的关系",
	}
}
