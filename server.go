package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"test.com/proto"
)

type echoService struct {
	proto.GrpcDemoServer
}

func (s echoService) GetEcho(ctx context.Context, message *proto.Echo) (*proto.Echo, error) {
	fmt.Printf("Echo ID: %s, message: %s \n", message.Id, message.Message)
	return message, nil
}

func main() {
	serverType := os.Getenv("SERVER_TYPE")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}
	server := echoService{}
	if serverType == "grpc" {
		grpcServer := grpc.NewServer()
		proto.RegisterGrpcDemoServer(grpcServer, server)
		fmt.Println("gRPC server listening on port 50051")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatal(err)
		}
		defer grpcServer.Stop()
	} else if serverType == "rest" {
		mux := runtime.NewServeMux()
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		err := proto.RegisterGrpcDemoHandlerServer(ctx, mux, server)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("REST server listening on port 50051")
		http.Serve(lis, mux)
	}
}
