package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"golang.org/x/net/context"
	// 导入grpc包
	"google.golang.org/grpc"
	// 导入刚才我们生成的代码所在的proto包。
	pb "test/protocol/admin"
)

func main() {
	// 连接grpc服务器
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()

	// 初始化Greeter服务客户端
	c := pb.NewAdminClient(conn)
	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 延迟关闭请求会话
	defer cancel()

	// 测试部分
	//rply, err := c.Registration(ctx, &pb.RegistrationReq{Name: "G1", Passwd: "mx179516"})
	//rply, err := c.Login(ctx, &pb.LoginReq{Name: "G1", Password: "mx179516"})
	//rply, err := c.Logout(ctx, &pb.LogoutReq{Name: "xiaowang"})

	// 这里后期需要注意 , 因为这里并不能返回error
	//rply, err := c.GetOnlineUserList(ctx, &pb.OnlineUserListReq{})
	//fmt.Println("rply: ", rply)
	//fmt.Println("rply.Name :", rply.Name)

	//rply, err := c.AddAndRemoveBlackList(ctx, &pb.AddAndRemoveBlackListReq{
	//	OwnerId:  2,
	//	TargetId: 3,
	//	Type:     3,
	//})

	// 发送信息的测试
	SendMessage_Fliter(c, ctx)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// 打印服务的返回的消息
	log.Printf("Greeting: %s", err)
}

// 这个写在客户端用于处理信息的过滤
func SendMessage_Fliter(c pb.AdminClient, ctx context.Context) {
	emptyJson, _ := json.Marshal(map[string]interface{}{})

	rply, _ := c.SendMessage(ctx, &pb.SendMessageReq{
		UserId:    2,
		RoomId:    1,
		IsSendAll: true,
		ToUserIds: string(emptyJson),
		Content:   "你好 , 世界",
		TargetId:  0,
	})
	fmt.Println("rply :", rply)
}
