package main

import (
	"context"
	"log"
	"net"
	"test/controller/login"
	"test/controller/logout"
	"test/controller/message"
	"test/controller/user"
	"test/controller/user_contact"
	"test/utils"
	"test/utils/database"

	pb "test/protocol/admin"

	// 注意v2版本
	"google.golang.org/grpc"
)

func init() {
	utils.InitConfig()
	database.InitMySQL()
	database.InitRedis()
}

type server struct {
	pb.UnimplementedAdminServer
}

func NewServer() pb.AdminServer {
	return &server{}
}

// Registration
// 用于用户的注册
func (s *server) Registration(ctx context.Context, in *pb.RegistrationReq) (*pb.RegistrationRsp, error) {
	err := user.AddUser(in.Name, in.Passwd)
	return nil, err
}

// Login
// 用于用户的登录
func (s *server) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRsp, error) {
	err := login.Login(in.Name, in.Password)
	return nil, err
}

// Logout
// 用于用户的登出
func (s *server) Logout(ctx context.Context, in *pb.LogoutReq) (*pb.LogoutRsp, error) {
	err := logout.Logout(in.Name)
	return nil, err
}

// 不足
// 这里的err 必然是nil 需要注意
// 感觉是因为 我GetUserList 返回引用的原因 , 等我先写完大致功能之后再回来改
func (s *server) GetOnlineUserList(ctx context.Context, in *pb.OnlineUserListReq) (out *pb.OnlineUserListRsp, err error) {
	//	先初始化再使用!
	out = &pb.OnlineUserListRsp{}
	out.Name, err = user.GetUserList()

	return out, err
}

// AddAndRemoveBlackList
// 黑名单操作

func (s *server) AddAndRemoveBlackList(ctx context.Context, in *pb.AddAndRemoveBlackListReq) (out *pb.AddAndRemoveBlackListRsp, err error) {
	// 因为是int32 这里做一层转换
	err = user_contact.ChangeFriendContact(int(in.OwnerId), int(in.TargetId), int(in.Type))

	return nil, err
}

// SendMessage
// 发送消息
func (s *server) SendMessage(ctx context.Context, in *pb.SendMessageReq) (out *pb.SendMessageRsp, err error) {
	message.AddMessage(in)
	return nil, err
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册Greeter service到server
	pb.RegisterAdminServer(s, &server{})
	// 启动gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:9999")
	log.Fatal(s.Serve(lis))
}
