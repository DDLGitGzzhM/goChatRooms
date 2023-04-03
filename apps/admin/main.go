package main

import (
	"context"
	"log"
	"net"
	"test/controller/user"
	"test/utils"

	pb "test/protocol/admin"

	// 注意v2版本
	"google.golang.org/grpc"
)

func init() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
}

type server struct {
	pb.UnimplementedAdminServer
}

func NewServer() pb.AdminServer {
	return &server{}
}

func (s *server) Registration(ctx context.Context, in *pb.RegistrationReq) (*pb.RegistrationReq, error) {
	err := user.AddUser(in.Name, in.Passwd)
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
