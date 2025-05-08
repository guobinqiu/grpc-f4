package main

import (
	"fmt"
	"io"
	"net"

	pb "github.com/guobinqiu/grpc-f4/bidirectional-streaming/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedChatServiceServer
}

func (s *server) Chat(stream pb.ChatService_ChatServer) error {
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		fmt.Printf("From User: %s, Text: %s\n", msg.User, msg.Text)

		// echo back
		msg.User = "server"
		stream.Send(msg)
	}
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &server{})
	fmt.Println("Server started on :50051")
	s.Serve(lis)
}
