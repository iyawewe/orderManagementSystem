package main

import (
	"net/http"
	pb "github.com/iyawewe/orderManagementSystem/common/api"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}
func (h *handler) registerRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /api/customers/{customersID}/orders", h.HandlerCreateOrder)
}

var items []*pb.ItemsWithQuantity
if err := common.ReadJSON(r,&items);err!=nil{
	common.WriteError(w,http.StatusBadRequest,err.Error())
	return
}

func (h *handler) HandlerCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	h.client.CreateOrder(r.Context(),&pb.CreateOrderRequest){
		CustomerId: customerID,
		Items: items,
	}
}
