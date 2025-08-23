package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "github.com/iyawewe/orderManagementSystem/common/api"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
	service OrdersServices
}

func NewGRPCHandler(grpcServer *grpc.Server, serservice OrdersServices) {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("new Order recieved! ORder %v", p)
	o := &pb.Order{
		Id: "42",
	}
	return o, nil
}
