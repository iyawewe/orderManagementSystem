package main

import (
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

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	})
}
