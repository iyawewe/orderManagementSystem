package main

import (
	"errors"
	"net/http"

	"github.com/iyawewe/orderManagementSystem/common/api"
	pb "github.com/iyawewe/orderManagementSystem/common/api"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandlerCreateOrder)
}

func (h *handler) HandlerCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var items []*pb.ItemsWithQuantity
	if err := api.ReadJSON(r, &items); err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	if err := validateItems(items); err != nil {
		api.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	o, err := h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
	if err != nil {
		api.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	api.WriteJSON(w, http.StatusOK, o)
}

func validateItems(items []*pb.ItemsWithQuantity) error {
	if len(items) == 0 {
		return errors.New("items must have at least one item")
	}
	for _, i := range items {
		if i.Id == "" {
			return errors.New("item ID is required")
		}
		if i.Quantity <= 0 {
			return errors.New("item quantity must be greater than 0")
		}
	}
	return nil
}
