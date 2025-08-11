package main

import "net/http"

type handler struct {
	//gateway
}

func NewHandler() *handler {
	return &handler{}
}
func (h *handler) registerRoutes(mux *http.ServeMux) {

	mux.HandleFunc("POST /api/customers/{customersID}/orders", h.HandlerCreateOrder)
}

func (h *handler) HandlerCreateOrder(w http.ResponseWriter, r *http.Request) {
}
