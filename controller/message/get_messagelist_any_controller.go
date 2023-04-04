package message

import (
	"test/dao"
	"test/models"
)

func GetMessageListAny() ([]*models.Message, error) {
	message := dao.GetMessageListAny()
	return message, MessageError{
		ErrorCode:    0,
		ErrorMessage: "查询成功",
	}
}
