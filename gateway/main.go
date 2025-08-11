package main

import (
	"log"
	"net/http"
)
common "github.com/iyawewew/orderManagementSystem/common"
var (
	httpAddr = common.EnvString("HTTP_ADDR",":8080")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)
	log.Printf("Starting server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("failed to start server")
	}
}
