package main

import (
	"fmt"
	"io"
	"net"
	"os"
	"test/client-streaming/proto"

	"google.golang.org/grpc"
)

type server struct {
	proto.UnimplementedFileServiceServer
}

func (s *server) Upload(stream proto.FileService_UploadServer) error {
	var file *os.File
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&proto.UploadStatus{
				Success: true,
				Message: "File uploaded successfully",
			})
		}
		if err != nil {
			return err
		}

		if file == nil {
			file, err = os.Create(chunk.Filename)
			if err != nil {
				return err
			}
		}
		file.Write(chunk.Content)
	}
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	proto.RegisterFileServiceServer(s, &server{})
	fmt.Println("Server started on :50051")
	s.Serve(lis)
}
