package main

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/guobinqiu/grpc-f4/client-streaming/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func upload(client proto.FileServiceClient, filename string) {
	file, _ := os.Open(filename)
	defer file.Close()

	// 客户端的 grpc Upload 方法
	stream, _ := client.Upload(context.Background())

	buf := make([]byte, 1024)
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

	// 文件上传成功后在server目录下生成test.txt
	upload(client, "test.txt")
}
