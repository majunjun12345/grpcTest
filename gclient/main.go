package main

import (
	"context"
	"fmt"
	"log"

	pb "grpcTest/gproto"

	"google.golang.org/grpc"
)

const (
	address = "127.0.0.1:1234"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("connection error:", err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)

	req := &pb.HelloRequest{}

	req.Name = "grpc"
	resp, err := client.SayHello(context.Background(), req)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Message)
}
