package main

import (
	"fmt"
	"io"
	"net"

	pb "github.com/guobinqiu/grpc-f4/bidirectional-streaming/proto"

	"google.golang.org/grpc"
)

/*
子类 由grpc自动生成
type UnimplementedChatServiceServer struct {
}
*/

// 父类 由我们自定义 开闭原则 对扩展开放、对修改关闭
type server struct {
	pb.UnimplementedChatServiceServer // 匿名嵌套 让父类拥有子类的方法和属性
}

/*
type server struct {
	inner pb.UnimplementedChatServiceServer // 具名嵌套
}
*/

/*
子类方法以placeholder的形式出现

	func (UnimplementedChatServiceServer) Chat(ChatService_ChatServer) error {
		return status.Errorf(codes.Unimplemented, "method Chat not implemented")
	}

我们要在父类里重写该方法
它是在客户端调用了客户端Chat方法后由grpc框架内部去调的它 你只要负责实现它就好
*/
func (s *server) Chat(stream pb.ChatService_ChatServer) error {
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
	pb.RegisterChatServiceServer(s, &server{})
	fmt.Println("Server started on :50051")
	s.Serve(lis)
}
