package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "ch14/path/to/your/package/hello"
	"google.golang.org/grpc"
)

func main() {
	// 连接服务器
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// 创建 gRPC 客户端
	client := pb.NewHelloServiceClient(conn)

	// 发送请求
	name := "World"
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}

	log.Printf("Greeting: %s", res.GetMessage())
}
