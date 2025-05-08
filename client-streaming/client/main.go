package main

import (
	"context"
	"fmt"
	"io"
	"os"
	"test/client-streaming/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func uploadFile(client proto.FileServiceClient, filePath string) {
	file, _ := os.Open(filePath)
	defer file.Close()

	stream, _ := client.Upload(context.Background())
	buf := make([]byte, 1024)
	filename := filePath

	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		stream.Send(&proto.FileChunk{
			Filename: filename,
			Content:  buf[:n],
		})
	}
	res, _ := stream.CloseAndRecv()
	fmt.Println("Upload:", res.Message)
}

func main() {
	conn, _ := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	client := proto.NewFileServiceClient(conn)

	uploadFile(client, "test.txt")
}
