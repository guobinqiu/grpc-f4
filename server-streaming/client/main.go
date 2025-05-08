package main

import (
	"context"
	"fmt"
	"io"
	"os"

	pb "github.com/guobinqiu/grpc-f4/server-streaming/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func download(client pb.FileServiceClient, filename string) {
	req := &pb.FileRequest{Filename: filename}

	// 客户端的 grpc Download 方法
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
	client := pb.NewFileServiceClient(conn)

	// 文件下载成功后在client目录下生成test.txt
	download(client, "test.txt")
}
