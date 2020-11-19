package main

import (
	"context"
	"fmt"
	helloService "go_learning/thirdparty/grpc/services/hello_service"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	t0 := time.Now()
    conn, err := grpc.Dial("localhost:10627", grpc.WithInsecure())
    if err != nil {
        log.Fatal(err)
	}
	t00 := time.Now()
	fmt.Printf("RPC Client Dial() cost: %v\n", t00.Sub(t0))
    defer conn.Close()

	for i := 0; i < 10; i++ {
		t1 := time.Now()
		client := helloService.NewHelloServiceClient(conn)
		reply, err := client.Hello(context.Background(), &helloService.String{Value: "zhang3"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("RPC Client Call Hello() return: " + reply.GetValue())
		t2 := time.Now()
		fmt.Printf("RPC Client Call Hello() cost: %v\n", t2.Sub(t1))
		
		reply, err = client.Hi(context.Background(), &helloService.String{Value: "li4"})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("RPC Client Call Hi() return: " + reply.GetValue())
		t3 := time.Now()
		fmt.Printf("RPC Client Call Hi() cost: %v\n", t3.Sub(t2))
	}
	
}