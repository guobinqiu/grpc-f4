package main

import (
	"context"
	"fmt"
	"io"
	"time"

	pb "github.com/guobinqiu/grpc-f4/bidirectional-streaming/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	// 客户端的 grpc Chat 方法
	stream, _ := client.Chat(context.Background())

	go func() {
		for _, txt := range []string{"hi", "how are you", "bye"} {
			stream.Send(&pb.ChatMessage{User: "client", Text: txt})
			time.Sleep(time.Second)
		}
		stream.CloseSend()
	}()

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Printf("From User: %s, Text: %s\n", msg.User, msg.Text)
	}
}
