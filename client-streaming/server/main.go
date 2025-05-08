package main

import (
	"fmt"
	"io"
	"net"
	"os"

	pb "github.com/guobinqiu/grpc-f4/client-streaming/proto"

	"google.golang.org/grpc"
)

/*
子类 由grpc自动生成
type UnimplementedFileServiceServer struct {
}
*/

// 父类 由我们自定义 开闭原则 对扩展开放、对修改关闭
type server struct {
	pb.UnimplementedFileServiceServer // 匿名嵌套 让父类拥有子类的方法和属性
}

/*
type server struct {
	inner pb.UnimplementedFileServiceServer // 具名嵌套
}
*/

/*
子类方法以placeholder的形式出现

	func (UnimplementedFileServiceServer) Upload(FileService_UploadServer) error {
		return status.Errorf(codes.Unimplemented, "method Upload not implemented")
	}

我们要在父类里重写该方法
它是在客户端调用了客户端Upload方法后由grpc框架内部去调的它 你只要负责实现它就好
*/
func (s *server) Upload(stream pb.FileService_UploadServer) error {
	var file *os.File
	defer func() {
		if file != nil {
			file.Close()
		}
	}()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.UploadStatus{
				Success: true,
				Message: "File successfully uploaded",
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
	pb.RegisterFileServiceServer(s, &server{})
	fmt.Println("Server started on :50051")
	s.Serve(lis)
}
