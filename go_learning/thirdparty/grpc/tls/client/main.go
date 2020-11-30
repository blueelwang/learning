package main

import (
	"context"
	"fmt"
	helloService "go_learning/thirdparty/grpc/services/hello_service"
	"io"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	t0 := time.Now()
	creds, err := credentials.NewClientTLSFromFile(
        "./tls/server.crt", "daemoncoder.com",
    )
    if err != nil {
        log.Fatal(err)
    }

    conn, err := grpc.Dial("10.221.81.129:10627", grpc.WithTransportCredentials(creds))
    if err != nil {
        log.Fatal(err)
    }
	t1 := time.Now()
	fmt.Printf("RPC Client Dial() cost: %v\n", t1.Sub(t0))
    defer conn.Close()

	for i := 0; i < 10; i++ {
		t2 := time.Now()
		client := helloService.NewHelloServiceClient(conn)
		t3 := time.Now()
		fmt.Printf("RPC Client Call NewHelloServiceClient() cost: %s\n", t3.Sub(t2))

		reply, err := client.Hello(context.Background(), &helloService.String{Value: "zhang3"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("RPC Client Call Hello() return: " + reply.GetValue())
		t4 := time.Now()
		fmt.Printf("RPC Client Call Hello() cost: %v\n", t4.Sub(t3))
		
		reply, err = client.Hi(context.Background(), &helloService.String{Value: "li4"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("RPC Client Call Hi() return: " + reply.GetValue())
		t5 := time.Now()
		fmt.Printf("RPC Client Call Hi() cost: %v\n", t5.Sub(t4))


		// stream
		stream, err := client.Channel(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		t6 := time.Now()
		fmt.Printf("RPC Client Call Channel() cost: %v\n", t6.Sub(t5))

		go func() {
			for j := 0; j < 3; j++ {
				t7 := time.Now()
				if err := stream.Send(&helloService.String{Value: "stream"}); err != nil {
					log.Fatal(err)
				}
				t8 := time.Now()
				fmt.Printf("RPC Client Call stream.Send() cost: %v\n", t8.Sub(t7))
			}
		}()
		for j := 0; j < 3; j++ {
			t9 := time.Now()
			reply, err := stream.Recv()
			t10 := time.Now()
			fmt.Printf("RPC Client Call stream.Recv() cost: %v\n", t10.Sub(t9))
			if err != nil {
				if err == io.EOF {
					break
				}
				log.Fatal(err)
			}
			fmt.Println(reply.GetValue())
		}
	}
}

func clientWithTLS() {
	
}
