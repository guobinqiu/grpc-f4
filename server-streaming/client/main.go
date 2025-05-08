package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"test/server-streaming/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func downloadFile(client proto.FileServiceClient, filename string) {
	req := &proto.FileRequest{Filename: filename}
	stream, _ := client.Download(context.Background(), req)

	file, _ := os.Create(filename)
	defer file.Close()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		file.Write(chunk.Content)
	}
	fmt.Println("Download complete")
}

func main() {
	conn, _ := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	client := proto.NewFileServiceClient(conn)

	downloadFile(client, "test.txt")
}
