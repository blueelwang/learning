package rpc

import (
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
	"log"

)

// ClientDemo _
func ClientDemo() {
	// client, err := rpc.Dial("tcp", "localhost:1234")
    conn, err := net.Dial("tcp", "localhost:1234")
    if err != nil {
        log.Fatal("dialing:", err)
    }

	client := rpc.NewClientWithCodec(jsonrpc.NewClientCodec(conn))

    var reply string
	err = client.Call("HelloService.Hello", "Zhang3", &reply)
	// jsonrpc实际发送的内容为：
	// {"method":"HelloService.Hello","params":["Zhang3"],"id":0}

    if err != nil {
        log.Fatal(err)
    }

    log.Println(reply)
}