package main

import (
	"fmt"
	"io"
	"net"
	"test/bidirectional-streaming/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedChatServiceServer
}

func (s *server) Chat(stream proto.ChatService_ChatServer) error {
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
	proto.RegisterChatServiceServer(s, &server{})
	fmt.Println("Server started on :50051")
	s.Serve(lis)
}
