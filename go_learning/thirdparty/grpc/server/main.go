package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	helloService "go_learning/thirdparty/grpc/services/hello_service"
)

func main() {
    grpcServer := grpc.NewServer()
    helloService.RegisterHelloServiceServer(grpcServer, new(helloService.HelloServiceImpl))

    lis, err := net.Listen("tcp", ":10627")
    if err != nil {
        log.Fatal(err)
    }
    grpcServer.Serve(lis)
}