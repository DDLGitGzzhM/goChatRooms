package message

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"test/dao"
	"test/models"
	pb "test/protocol/admin"
	"time"
)

func AddMessageNotAll(req *pb.SendMessageReq) MessageError {
	//处理toUserIds
	//然后进行单一的发送
	Ids := strings.Split(req.ToUserIds, ",")

	//fmt.Println("req.ToUserIds", req.ToUserIds)
	toUserIds_json, err := json.Marshal(Ids)

	if err != nil {
		fmt.Println(err.Error())
		return MessageError{
			ErrorCode:    1,
			ErrorMessage: "AddMessageNotall error",
		}
	}
	for _, val := range Ids {
		t_val, _ := strconv.Atoi(val)
		message_mysql := models.Message{
			UserId:    int(req.UserId),
			RoomId:    int(req.RoomId),
			IsSendAll: false,
			ToUserIds: string(toUserIds_json),
			Content:   req.Content,
			TargetId:  t_val,
			CTime:     time.Now(),
			MTime:     time.Now(),
		}
		dao.AddMessage_Mysql(message_mysql)
	}
	return MessageError{
		ErrorCode:    0,
		ErrorMessage: "发送成功",
	}
}
