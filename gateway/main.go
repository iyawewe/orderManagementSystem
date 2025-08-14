package main

import (
	"log"
	"net/http"

	"github.com/iyawewe/orderManagementSystem/common"
	pb "github.com/iyawewe/orderManagementSystem/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr = "localhost:2000"
)

func main() {

	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to dial service: %v", err)
	}
	defer conn.Close()

	log.Println("dialing orders services at", orderServiceAddr)
	c := pb.NewOrderServiceClient(conn)
	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)
	log.Printf("Starting server on %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log.Fatalf("failed to start server")
	}
}
