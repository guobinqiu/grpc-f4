package main

import (
	"fmt"
	"io"
	"net"
	"os"

	pb "github.com/guobinqiu/grpc-f4/server-streaming/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServiceServer
}

func (s *server) Download(req *pb.FileRequest, stream pb.FileService_DownloadServer) error {
	file, err := os.Open(req.Filename)
	if err != nil {
		return err
	}
	defer file.Close()

	buf := make([]byte, 1024)
	for {
		n, err := file.Read(buf)
		if err == io.EOF {
			break
		}
		stream.Send(&pb.FileChunk{Content: buf[:n]})
	}
	return nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, &server{})
	fmt.Println("Server started on :50051")
	s.Serve(lis)
}
