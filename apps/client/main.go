package main

import (
	"log"
	"time"

	"golang.org/x/net/context"
	// 导入grpc包
	"google.golang.org/grpc"
	// 导入刚才我们生成的代码所在的proto包。
	pb "test/protocol/helloworld"
)

const (
	defaultName = "world"
)

func main() {
	// 连接grpc服务器
	conn, err := grpc.Dial("localhost:8090", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 延迟关闭连接
	defer conn.Close()

	// 初始化Greeter服务客户端
	c := pb.NewGreeterClient(conn)

	// 初始化上下文，设置请求超时时间为1秒
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	// 延迟关闭请求会话
	defer cancel()

	// 调用SayHello接口，发送一条消息
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	// 打印服务的返回的消息
	log.Printf("Greeting: %s", r.Message)
}
