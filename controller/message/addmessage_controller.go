package message

import (
	"fmt"
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

func AddMessage(req *pb.SendMessageReq) MessageError {
	message := models.Message{
		UserId:    int(req.UserId),
		RoomId:    int(req.RoomId),
		IsSendAll: req.IsSendAll,
		ToUserIds: req.ToUserIds,
		Content:   req.Content,
		CTime:     time.Now(),
		MTime:     time.Now(),
	}
	dao.AddMessage(message)
	return MessageError{
		ErrorCode:    0,
		ErrorMessage: "发送成功",
	}
}
