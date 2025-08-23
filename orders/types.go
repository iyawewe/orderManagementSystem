package main

import (
	"context"

	pb "github.com/iyawewe/orderManagementSystem/common/api"
)

type OrdersServices interface {
	CreateOrder(context.Context) error
	validateOrder(context.Context, *pb.CreateOrderRequest) error
}
type OrdersStore interface {
	Create(context.Context) error
}
