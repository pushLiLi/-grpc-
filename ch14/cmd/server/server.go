package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "ch14/path/to/your/package/hello"
	"google.golang.org/grpc"
)

// 定义服务的实现
type helloServer struct {
	pb.UnimplementedHelloServiceServer
}

func (s *helloServer) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	message := fmt.Sprintf("Hello, %s!", req.Name)
	return &pb.HelloResponse{Message: message}, nil
}

func main() {
	// 监听本地端口
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// 创建新的 gRPC 服务器
	grpcServer := grpc.NewServer()

	// 注册服务
	pb.RegisterHelloServiceServer(grpcServer, &helloServer{})

	log.Printf("Server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
