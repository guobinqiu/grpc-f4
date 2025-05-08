package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/guobinqiu/grpc-f4/simple/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetingServiceServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + req.Name}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})
	fmt.Println("Server started on :50051")
	s.Serve(lis)
}
