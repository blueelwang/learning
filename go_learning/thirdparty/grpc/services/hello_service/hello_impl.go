package hello_service

import (
	"fmt"
	"context"
)

type HelloServiceImpl struct{}

func (p *HelloServiceImpl) Hello(ctx context.Context, args *String) (*String, error) {
	fmt.Println("RPC Server Hello() received content: " + args.GetValue())
    reply := &String{Value: "hello:" + args.GetValue()}
    return reply, nil
}

func (p *HelloServiceImpl) Hi(ctx context.Context, args *String) (*String, error) {
	fmt.Println("RPC Server Hi() received content: " + args.GetValue())
    reply := &String{Value: "hi:" + args.GetValue()}
    return reply, nil
}

func (p *HelloServiceImpl) mustEmbedUnimplementedHelloServiceServer() {}
