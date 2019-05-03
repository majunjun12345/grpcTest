package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
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
	// 创建对象
	arith := new(Arith)
	// 注册对象
	rpc.Register(arith)
	//指定rpc的传输协议 这里采用http协议作为rpc调用的载体 也可以用rpc.ServeConn处理单个连接请求
	rpc.HandleHTTP()

	/*
		one
	*/
	err := http.ListenAndServe("127.0.0.1:1234", nil) // 加上返回值 err 后就阻塞了
	if err != nil {
		log.Fatal("err:", err)
	}
	// log.Fatalln("err:", http.ListenAndServe("127.0.0.1:1234", nil))

	/*
		two
	*/
	// listen, err := net.Listen("tcp", ":1234")
	// if err != nil {
	// 	log.Fatal("listrening err:", err)
	// }
	// go http.Serve(listen, nil)
	// os.Stdin.Read(make([]byte, 1)) // 这个也是起到阻塞的作用

	/*
		three
	*/
	// listen, err := net.Listen("tcp", "127.0.0.1:1234")
	// if err != nil {
	// 	log.Fatalln("fatal err:", err)
	// }
	// for {
	// 	conn, err := listen.Accept()

	// 	if err != nil {
	// 		continue
	// 	}

	// 	go func(conn net.Conn) {
	// 		fmt.Fprintf(os.Stdout, "%s", "new connection coming...\n")
	// 		rpc.ServeConn(conn)

	// 	}(conn)
	// }
}

/*
	server:
	定义一个对象
	对象的方法必须是可导出的
	方法接收两个参数，一个是请求参数，一个是响应参数，响应参数必须是指针类型；
	请求结构体和响应结构体字段也必须是可导出的；
	返回一个 error

	创建新对象，必须用 new，返回一个指针
	将对象注册进 rpc
	开启 http 或 tcp 服务

	client:
	client 端构造与 server 一致的请求响应结构体
	建立 tcp 或 http 连接
	通过连接进行 rpc 调用，第一个参数是对象的方法(字符串形式)，第二三个参数分别是请求响应参数，响应参数是指针形式

	http 和 tcp：
	前两种都是基于 http 的 rpc 调用，客户端使用 rpc.DialHTTP 即可
	第三种属于基于 tcp rpc 调用，需要阻塞不断接收连接，客户端使用 rpc.Dial

	jsonrpc：
	使用 jsonrpc 将 rpc 换成 jsonrpc 即可
	rpc 默认使用 gob 编码，jsonrpc 默认使用 json 编码

	同步异步：
	同步异步是针对客户端而言的，同步调用直到有结果返回才能执行下一步
	异步发出调用信号后就可以做其他事情了，可以通过 chan 获取返回结果
*/
