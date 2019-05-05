package main

import (
	"fmt"
	"log"
	"net/rpc"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, rem int
}

func main() {

	// tcp 使用 rpc.Dial
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:1234")
	if err != nil {
		log.Fatal("dial err:", err)
	}

	args := Args{10, 8}
	var resp int

	// call 实现 rpc 的同步调用
	err = client.Call("Arith.Multiply", args, &resp)
	if err != nil {
		log.Fatal("call Multiply err:", err)
	}
	log.Printf("Arith %d * %d=%d", args.A, args.B, resp)

	q := Quotient{}
	err = client.Call("Arith.Divide", args, &q)
	if err != nil {
		log.Fatal("call Divide err:", err)
	}
	log.Printf("Arith %d * %d=%#v", args.A, args.B, q)

	// go 实现 rpc 异步调用
	/*

		cient.Go的返回值包含了最后一个参数(channel),想获取调用结果，可以从参数管道中直接获取，也可以从返回值Done中获取;
		参数中的 chan 可以是 nil

	*/
	done := client.Go("Arith.Divide", args, &q, nil)
	for {
		select {
		case <-done.Done:
			log.Printf("Arith %d * %d=%#v", args.A, args.B, q)
			goto loop
		default:
			fmt.Println("继续向下执行....")
			time.Sleep(time.Second * 1)
		}
	}
loop:
}
