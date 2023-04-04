package message

import (
	"test/dao"
	"test/models"
)

func GetMessageListNotAny(OwnerId, TargetId int) ([]*models.Message, error) {
	message := dao.GetMessageListNotAny(OwnerId, TargetId)
	return message, MessageError{
		ErrorCode:    0,
		ErrorMessage: "查询成功",
	}
}
