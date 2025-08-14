package main

import (
	"context"
	"log"
	"net"

	"github.com/iyawewe/orderManagementSystem/common"
	"google.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	grpcServer := grpc.NewServer()

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to connect to grpc server: %v", err)
	}
	defer l.Close()

	store := NewStrore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer)

	svc.CreateOrder(context.Background())
	log.Printf("gRPC server listening on %s", grpcAddr)
	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}

}
