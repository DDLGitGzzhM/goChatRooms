package message

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"test/dao"
	"test/models"
	pb "test/protocol/admin"
	"time"
)

// 用于 返回 用户CRUD时候的错误
// 进行用户的格式规范
type MessageError struct {
	//  错误码
	ErrorCode int
	//	错误信息
	ErrorMessage string
}

func (e MessageError) Error() string {
	return fmt.Sprintf("error code: %d, message: %s", e.ErrorCode, e.ErrorMessage)
}

// 这个是存放在Mysql里面做持久化
func AddMessage_Mysql(req *pb.SendMessageReq) MessageError {
	message := models.Message{
		UserId:    int(req.UserId),
		RoomId:    int(req.RoomId),
		IsSendAll: req.IsSendAll,
		ToUserIds: req.ToUserIds,
		Content:   req.Content,
		CTime:     time.Now(),
		MTime:     time.Now(),
	}
	dao.AddMessage_Mysql(message)
	return MessageError{
		ErrorCode:    0,
		ErrorMessage: "发送成功",
	}
}

// 正常的流程应该是将其放入 Redis
// 然后Mysql在从Redis中放入存储
func AddMessage(req *pb.SendMessageReq) MessageError {
	message := map[string]interface{}{
		"UserId":    int(req.UserId),
		"RoomId":    int(req.RoomId),
		"IsSendAll": req.IsSendAll,
		"ToUserIds": req.ToUserIds,
		"Content":   req.Content,
		"TargetId":  req.TargetId,
		"CTime":     time.Now(),
		"MTime":     time.Now(),
	}

	message_json, err := json.Marshal(message)
	if err != nil {
		log.Fatalln("message_json err :", message_json)
		return MessageError{
			ErrorCode:    1,
			ErrorMessage: "message_json 初始化失败",
		}
	}

	roomId := strconv.Itoa(int(req.RoomId))
	dao.AddMessage(int(time.Now().Unix()), roomId, string(message_json))

	/*
		这里需要考虑何时存放进mysql里面
		目前有的做法考虑是搞一个全局变量计数器 , 当计数器到达一定的数量的时候 我们持久化进数据库
		不过这里选择 , 存一个放一个
	*/

	message_mysql := models.Message{
		UserId:    int(req.UserId),
		RoomId:    int(req.RoomId),
		IsSendAll: req.IsSendAll,
		ToUserIds: req.ToUserIds,
		Content:   req.Content,
		TargetId:  int(req.TargetId),
		CTime:     time.Now(),
		MTime:     time.Now(),
	}
	dao.AddMessage_Mysql(message_mysql)

	return MessageError{
		ErrorCode:    0,
		ErrorMessage: "发送成功",
	}
}
