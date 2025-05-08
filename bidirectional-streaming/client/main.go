package main

import (
	"context"
	"fmt"
	"io"
	"test/bidirectional-streaming/proto"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	client := proto.NewChatServiceClient(conn)

	stream, _ := client.Chat(context.Background())

	go func() {
		for _, txt := range []string{"hi", "how are you", "bye"} {
			stream.Send(&proto.ChatMessage{User: "client", Text: txt})
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
