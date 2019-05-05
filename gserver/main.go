package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"

	// 定位到文件夹即可，文件夹和包名可以不同
	pb "grpcTest/gproto"
)

const (
	address = "127.0.0.1:1234"
)

type helloService struct {
}

func (h helloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	resp := &pb.HelloReply{} // resp := new(pb.HelloReply)
	resp.Message = "hello" + req.Name + "."
	return resp, nil
}

func main() {

	helloServer := helloService{} // var helloServer helloService   helloServer := new(helloService)
	// 实现gRPC Server
	s := grpc.NewServer()
	// 注册helloServer为客户端提供服务
	pb.RegisterHelloServer(s, helloServer)

	// listen, err := net.Listen("tcp", address)
	// if err != nil {
	// 	log.Fatal("err:%s\n", err)
	// }
	// s.Serve(listen)

	fmt.Println("Listen on" + address)
	log.Fatal(http.ListenAndServe(address, nil))
}

/*
	总结：
	proto 定义请求和响应结构体，定义暴露的方法，方法的参数是请求结构体，返回响应结构体；(我们可以一次性的在一个 .proto 文件中定义服务并使用任何支持它的语言去实现客户端和服务器)
	在 server 中调用 RegisterHelloServer，在 client 调用 NewHelloClient

	server：
	定义一个对象，实现 proto 文件中暴露的方法
	在 main 中实例化这个类，并将这个类注册进实例话的 grpc server
	提供服务即可

	client：
	通过 grpc 和 生成的 proto 文件实例话 client
	构造并实例化请求结构体，通过实例话的 client 发出请求即可！

*/
