package hello_service

import (
	"fmt"
	"context"
	"io"
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

func (p *HelloServiceImpl) Channel(stream HelloService_ChannelServer) error {
	for {
		args, err := stream.Recv()
		if err != nil {
			if err == io.EOF {
				return nil
			}
			return err
		}
		fmt.Println("RPC Server Channel() received content: " + args.GetValue())
		reply := &String{Value: "hello:" + args.GetValue()}

		err = stream.Send(reply)
		if err != nil {
			return err
		}
	}
}

func (p *HelloServiceImpl) mustEmbedUnimplementedHelloServiceServer() {}
