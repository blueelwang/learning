package systemcall

import (
	"fmt"
	"net"
	"time"
)

func SocketDemo() {
	
	done := make(chan struct{}, 1)
	go SocketServerDemo(done)
	time.Sleep(1 * time.Second)

	go SocketClientDemo()
	go SocketClientDemo()
	go SocketClientDemo()
	go SocketClientDemo()
	go SocketClientDemo()

	<-done

}

func SocketServerDemo(done chan struct{}) {
	listener, _ := net.Listen("tcp4", "127.0.0.1:8088")
	
	for i := 0; i < 5; i++ {
		conn, _ := listener.Accept()
		buffer := make([]byte, 1024)
		conn.Read(buffer)
		fmt.Printf("%s\n", buffer)
		conn.Close()
	}
	listener.Close()
	done <- struct{}{}
}

func SocketClientDemo() {
	// conn, _ := net.Dial("tcp4", "127.0.0.1:8088")
	conn, _ := net.DialTimeout("tcp4", "127.0.0.1:8088", 2 * time.Second)
	data := "www.daemoncoder.com"
	conn.Write(([]byte)(data))
	conn.Close()
}

/*
Go socket 底层使用的是非阻塞模式，但是 Go 的 socket api为我们屏蔽了非阻塞读的细节处理，像部分写等。
*/
