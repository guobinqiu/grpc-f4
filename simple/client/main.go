package main

import (
	"context"
	"fmt"

	pb "github.com/guobinqiu/grpc-f4/simple/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()

	client := pb.NewGreetingServiceClient(conn)

	res, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: "Guobin"})
	fmt.Println(res.Message)
}
