package message

import (
	"encoding/json"
	"strconv"
	"test/dao"
	"test/models"
)

func GetMessageListAny(roomId int) ([]string, []string, error) {
	troomId := strconv.Itoa(roomId)
	temp := dao.GetMessageListAny(troomId)

	message := make([]*models.Message, 0, 50)

	for _, val := range temp {
		temp_message := &models.Message{}
		err := json.Unmarshal([]byte(val), &temp_message)
		if err != nil {
			return nil, nil, MessageError{
				ErrorCode:    1,
				ErrorMessage: "获取信息失败",
			}
		}
		message = append(message, temp_message)
	}

	content := make([]string, 0, 50)
	time := make([]string, 0, 50)
	//for()
	return content, time, MessageError{
		ErrorCode:    0,
		ErrorMessage: "查询成功",
	}
}
