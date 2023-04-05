package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
	pb "test/protocol/admin"
)

func main() {
	// 创建 gin 引擎
	r := gin.Default()

	// 连接 gRPC 服务
	conn, err := grpc.Dial("localhost:9999"+
		"", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed to dial: %v", err)
	}

	// 创建反向代理
	mux := runtime.NewServeMux()

	// 注册反向代理
	err = pb.RegisterAdminHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalf("failed to register gateway: %v", err)
	}

	// 注册路由
	r.GET("/register", func(c *gin.Context) {
		//// 从 URL 参数中获取 name
		name := c.Query("name")
		pasword := c.Query("password")

		// 调用 gRPC 服务
		req := &pb.RegistrationReq{Name: name, Passwd: pasword}
		_, err := pb.NewAdminClient(conn).Registration(context.Background(), req)
		fmt.Println("err >>>>>", err)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error()})
	})
	r.GET("/sendMessage", func(c *gin.Context) {

		content := c.Query("message")

		emptyJson, _ := json.Marshal(map[string]interface{}{})
		rply, _ := pb.NewAdminClient(conn).SendMessage(context.Background(), &pb.SendMessageReq{
			UserId:    2,
			RoomId:    1,
			IsSendAll: true,
			ToUserIds: string(emptyJson),
			Content:   content,
			TargetId:  0,
		})
		fmt.Println("rply :", rply)

	})

	r.GET("/getMessage", func(c *gin.Context) {

		Type := c.Query("Type")
		roomId := c.Query("roomId")
		TargetId := c.Query("TargetId")

		tType, _ := strconv.Atoi(Type)
		tRoomId, _ := strconv.Atoi(roomId)
		tTargetId, _ := strconv.Atoi(TargetId)
		rply, _ := pb.NewAdminClient(conn).GetMessage(context.Background(), &pb.GetMessageReq{
			Type:     int64(tType),
			RoomId:   int64(tRoomId),
			TargetId: int64(tTargetId),
		})
		fmt.Println("rply :", rply)
		c.JSON(http.StatusOK, gin.H{
			"message": rply})
	})

	// 启动 HTTP 服务器
	err = r.Run(":9391")
	if err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
