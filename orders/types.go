package main

import "context"

type OrdersServices interface {
	CreateOrder(context.Context) error
}
type OrdersStore interface {
	Create(context.Context) error
}
