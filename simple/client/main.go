package main

import (
	"context"
	"fmt"
	"test/simple/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	defer conn.Close()
	client := proto.NewGreetServiceClient(conn)

	res, _ := client.SayHello(context.Background(), &proto.HelloRequest{Name: "Guobin"})
	fmt.Println(res.Message)
}
