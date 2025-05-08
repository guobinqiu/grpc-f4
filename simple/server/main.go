package main

import (
	"context"
	"fmt"
	"net"

	pb "github.com/guobinqiu/grpc-f4/simple/proto"

	"google.golang.org/grpc"
)

/*
子类 由grpc自动生成
type UnimplementedGreetingServiceServer struct{}
*/

// 父类 由我们自定义 开闭原则 对扩展开放、对修改关闭
type server struct {
	pb.UnimplementedGreetingServiceServer // 匿名嵌套 让父类拥有子类的方法和属性
}

/*
type server struct {
	inner pb.UnimplementedGreetingServiceServer // 具名嵌套
}
*/

/*
子类方法以placeholder的形式出现

	func (UnimplementedGreetingServiceServer) SayHello(context.Context, *HelloRequest) (*HelloReply, error) {
		return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
	}

我们要在父类里重写该方法
它是在客户端调用了客户端SayHello方法后由grpc框架内部去调的它 你只要负责实现它就好
*/
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello, " + req.Name}, nil
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	s := grpc.NewServer()
	pb.RegisterGreetingServiceServer(s, &server{})
	fmt.Println("Server started on :50051")
	s.Serve(lis)
}
