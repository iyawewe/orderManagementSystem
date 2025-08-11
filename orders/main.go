package main

import "context"

func main() {
	store := NewStrore()
	svc := NewService(store)

	svc.CreateOrder(context.Background())

}
