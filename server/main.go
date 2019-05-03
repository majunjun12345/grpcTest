package main

import (
	"errors"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"time"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, resp *int) error {
	time.Sleep(3 * time.Second)
	*resp = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, resp *Quotient) error {
	time.Sleep(3 * time.Second)
	if args.B == 0 {
		return errors.New("divide by zero!")
	}
	resp.Quo = args.A / args.B
	resp.Rem = args.A % args.B
	return nil
}

func main() {
	arith := new(Arith)
	rpc.Register(arith)
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listrening err:", err)
	}
	go http.Serve(listen, nil)
	os.Stdin.Read(make([]byte, 1))
}
