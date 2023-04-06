package message

import (
	"encoding/json"
	"strconv"
	"test/dao"
	"test/models"
)

// http://localhost:9391/getMessage?Type=1&roomId=1&TargetId=0
// http://localhost:9391/sendMessage?message=消息,消息
func GetMessageListAny(roomId int) ([]string, []string, []string, error) {
	troomId := strconv.Itoa(roomId)
	temp := dao.GetMessageListAny(troomId)

	message := make([]*models.Message, 0, 50)

	for _, val := range temp {
		temp_message := &models.Message{}
		err := json.Unmarshal([]byte(val), &temp_message)
		if err != nil {
			return nil, nil, nil, MessageError{
				ErrorCode:    1,
				ErrorMessage: "获取信息失败",
			}
		}
		message = append(message, temp_message)
	}

	content := make([]string, 0, 50)
	time := make([]string, 0, 50)
	targetId := make([]string, 0, 50)

	for _, val := range message {
		content = append(content, val.Content)
		time = append(time, (val.CTime.String()))
		targetId = append(targetId, strconv.Itoa(val.UserId))
		//fmt.Println("content :", content)
		//fmt.Println("time : ", val.CTime.String())
		//fmt.Println("TargetId :", targetId)
	}

	return content, time, targetId, nil
	//MessageError{
	//	ErrorCode:    0,
	//	ErrorMessage: "查询成功",
	//}
}
