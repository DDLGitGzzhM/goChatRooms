package dao

import (
	"context"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"test/models"
	"test/utils/database"
)

// 那么这里采用 发一次就存一次的形式咯
func AddMessage_Mysql(message models.Message) *gorm.DB {
	return database.DB.Create(&message)
}

// key : RoomId
// value : message_json ( .... )
// message_json[content : "消息内容]
func AddMessage(timestamp int, roomId, message_json string) {
	database.Red.ZAdd(context.Background(), roomId, &redis.Z{
		Score:  float64(timestamp),
		Member: message_json,
	})
}

func GetMessageListAny(roomId string) []string {
	message_jsons, err := database.Red.ZRevRange(context.Background(), roomId, 0, 49).Result()
	if err != nil {
		return nil
	}

	return message_jsons
}

// 首先 问题是 toUserIds是一个json
// 懂了 toUserIds 其实是一个冗余的 变量
// 所以 我们需要引入 TargetId
func GetMessageListNotAny(OwnerId, TargetId int) []*models.Message {
	message := make([]*models.Message, 0, 50)
	database.DB.Where("userId = ? and targetId = ?", OwnerId, TargetId).Find(&message)
	return message
}
