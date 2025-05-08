package main

import (
	"context"
	"fmt"
	"net"
	"test/simple/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedGreetServiceServer
}

func (s *server) SayHello(ctx context.Context, req *proto.HelloRequest) (*proto.HelloReply, error) {
	return &proto.HelloReply{Message: "Hello, " + req.Name}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	proto.RegisterGreetServiceServer(s, &server{})
	fmt.Println("Server started on :50051")
	s.Serve(lis)
}
