package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	helloService "go_learning/thirdparty/grpc/services/hello_service"
)

func main() {
	creds, err := credentials.NewServerTLSFromFile("./tls/server.crt", "./tls/server.key")
    if err != nil {
        log.Fatal(err)
    }

	grpcServer := grpc.NewServer(grpc.Creds(creds))
	helloService.RegisterHelloServiceServer(grpcServer, new(helloService.HelloServiceImpl))

    lis, err := net.Listen("tcp", ":10627")
    if err != nil {
        log.Fatal(err)
    }
    grpcServer.Serve(lis)
}