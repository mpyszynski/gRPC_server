package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"test.com/proto"
)

type echoService struct {
	proto.GrpcDemoServer
}

func (s echoService) GetEcho(ctx context.Context, in *proto.Echo) (*proto.Echo, error) {
	fmt.Printf("Echo ID: %s, message: %s \n", in.Id, in.Message)
	return in, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	server := echoService{}
	proto.RegisterGrpcDemoServer(grpcServer, server)
	fmt.Println("gRPC server listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatal(err)
	}
	defer grpcServer.Stop()
}
