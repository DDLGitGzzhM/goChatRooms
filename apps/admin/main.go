package main

import (
	"context"
	"log"
	"net"

	pb "test/protocol/admin"

	// 注意v2版本
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedAdminServer
}

func NewServer() pb.AdminServer {
	return &server{}
}

func (s *server) Registration(context.Context, *pb.RegistrationReq) (*pb.RegistrationReq, error) {
	return nil, nil
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// 创建一个gRPC server对象
	s := grpc.NewServer()
	// 注册Greeter service到server
	pb.RegisterAdminServer(s, &server{})
	// 启动gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
