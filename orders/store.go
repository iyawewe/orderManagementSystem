package main

import "context"

type store struct {
	//mongodb in future
}

// Create implements OrdersStore.
func (s *store) Create(context.Context) error {
	panic("unimplemented")
}

// CreateOrder implements OrdersServices.
func (s *store) CreateOrder(context.Context) error {
	panic("unimplemented")
}

func NewStrore() *store {
	return &store{}
}
func (s *store) create(context.Context) error {
	// logic to create order in store
	return nil
}
