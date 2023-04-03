package main

import (
	"context"
	"log"
	"net"
	"test/controller/login"
	"test/controller/logout"
	"test/controller/user"
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

func (s *server) Registration(ctx context.Context, in *pb.RegistrationReq) (*pb.RegistrationRsp, error) {
	err := user.AddUser(in.Name, in.Passwd)
	return nil, err
}
func (s *server) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRsp, error) {
	err := login.Login(in.Name, in.Password)
	return nil, err
}
func (s *server) Logout(ctx context.Context, in *pb.LogoutReq) (*pb.LogoutRsp, error) {
	err := logout.Logout(in.Name)
	return nil, err
}
func (s *server) GetOnlineUserList(ctx context.Context, in *pb.OnlineUserListReq) (out *pb.OnlineUserListRsp, err error) {
	out.Name, err = user.GetUserList()
	return out, err
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
