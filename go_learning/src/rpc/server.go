package rpc

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"log"
)

// HelloService _
type HelloService struct {}

// Hello Go语言的RPC规则：方法只能有两个可序列化的参数，其中第二个参数是指针类型，并且返回一个error类型，同时必须是公开的方法。
func (p *HelloService) Hello(request string, reply *string) error {
    *reply = "hello," + request
    return nil
}

// ServerDemo _
func ServerDemo() {
	rpc.RegisterName("HelloService", new(HelloService))

    listener, err := net.Listen("tcp", ":1234")
    if err != nil {
        log.Fatal("ListenTCP error:", err)
    }

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}
	
		// go rpc.ServeConn(conn)
		go rpc.ServeCodec(jsonrpc.NewServerCodec(conn))
	}

}