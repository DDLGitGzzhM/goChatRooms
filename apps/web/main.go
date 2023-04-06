package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"strconv"
	pb "test/protocol/admin"
)

/*
所有路由均采用get的方式 ,方便进行调试

注册 /register
http://localhost:9391/register?name=G4&password=mx179516

登录 /login
http://localhost:9391/login?name=G4&password=mx179516

下线 /logout
http://localhost:9391/logout?name=G2&password=mx179516

注销 /logff
http://localhost:9391/logoff?name=G1&password=mx179516

获得用户在线和离线人数 /GetUserList
http://localhost:9391/getUserList

拉黑成功 /AddAndRemoveBlackList
http://localhost:9391/addAndRemoveBlackList?ownerId=5&targetId=6&type=2

发送消息操作 /SendMessage
发送给全部人
http://localhost:9391/sendMessage?userId=4&roomId=1&type=true&content=%22test_1%22&toUserIds={}
发送给单独的两个人 5 发给 6，7
http://localhost:9391/sendMessage?userId=5&roomId=1&type=false&content=%22test_1%22&toUserIds=[%22{\%22id1\%22:\%226\%22,\%22id2\%22:\%227\%22}%22]
*/
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

		c.JSON(http.StatusOK, gin.H{
			"message": err.Error()})
	})
	//登录路由
	r.GET("/login", func(c *gin.Context) {
		//获取登录要用的参数
		name := c.Query("name")
		pasword := c.Query("password")

		//调用gRPC服务
		req := &pb.LoginReq{Name: name, Password: pasword}
		_, err := pb.NewAdminClient(conn).Login(context.Background(), req)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error()})
	})
	//用户下线的路由
	r.GET("/logout", func(c *gin.Context) {
		//获取下线要用的参数
		name := c.Query("name")

		//调用gRPC服务
		req := &pb.LogoutReq{Name: name}
		_, err := pb.NewAdminClient(conn).Logout(context.Background(), req)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error()})
	})

	//用户的注销
	r.GET("/logoff", func(c *gin.Context) {
		//获取注销需要的参数
		name := c.Query("name")

		//调用gRPC服务
		req := &pb.LogoffReq{Name: name}
		_, err := pb.NewAdminClient(conn).Logoff(context.Background(), req)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error()})
	})
	//获取在线和离线的人数
	r.GET("/getUserList", func(c *gin.Context) {
		req := &pb.UserListReq{}
		resp, _ := pb.NewAdminClient(conn).GetUserList(context.Background(), req)
		c.JSON(http.StatusOK, gin.H{
			"messgae": resp,
		})
	})
	//加入黑名单
	r.GET("/addAndRemoveBlackList", func(c *gin.Context) {
		ownerId, _ := strconv.Atoi(c.Query("ownerId"))
		targetId, _ := strconv.Atoi(c.Query("targetId"))
		_type, _ := strconv.Atoi(c.Query("type"))
		//1表示加白
		//2表示拉黑
		req := &pb.AddAndRemoveBlackListReq{
			OwnerId:  int32(ownerId),
			TargetId: int32(targetId),
			Type:     int32(_type),
		}
		_, err := pb.NewAdminClient(conn).AddAndRemoveBlackList(context.Background(), req)
		c.JSON(http.StatusOK, gin.H{
			"message": err.Error(),
		})
	})
	//发送消息
	r.GET("/sendMessage", func(c *gin.Context) {
		userId, _ := strconv.Atoi(c.Query("userId"))
		roomId, _ := strconv.Atoi(c.Query("roomId"))
		toUserIds := c.Query("toUserIds")
		bool_type, _ := strconv.ParseBool(c.Query("type"))
		content := c.Query("content")

		//emptyJson, _ := json.Marshal(map[string]interface{}{})
		fmt.Println("bool_type :", bool_type)
		_, err := pb.NewAdminClient(conn).SendMessage(context.Background(), &pb.SendMessageReq{
			UserId:    int64(userId),
			RoomId:    int64(roomId),
			IsSendAll: bool_type,
			ToUserIds: toUserIds,
			Content:   content,
		})
		c.JSON(http.StatusOK, gin.H{
			"messsage": err.Error(),
		})

	})

	//获得消息
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
