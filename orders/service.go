package main

import (
	"context"
	"log"

	"github.com/iyawewe/orderManagementSystem/common/api"
	pb "github.com/iyawewe/orderManagementSystem/common/api"
)

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{store}
}
func (s *service) CreateOrder(context.Context) error {
	return nil
}

func (s *service) validateOrder(ctx context.Context, p *pb.CreateOrderRequest) error {
	if len(p.Items) == 0 {
		return api.ErrNoItems
	}
	mergedItems := mergedItemsQuantities(p.Items)
	log.Print(mergedItems)
	//validaet ewotht the stock servoce
	return nil
}

func mergedItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {
	merged := make([]*pb.ItemsWithQuantity, 0)

	for _, item := range items {
		found := false
		for _, finalItem := range merged {
			if finalItem.Id == item.Id {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}
		if !found {
			merged = append(merged, item)
		}
	}
	return merged
}
