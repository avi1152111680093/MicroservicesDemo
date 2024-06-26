package main

import (
	"fmt"
	"log"

	// "github.com/avi1152111680093/microservices/order/config"
	"github.com/avi1152111680093/microservices/order/internal/adapters/db"
	"github.com/avi1152111680093/microservices/order/internal/adapters/grpc"
	"github.com/avi1152111680093/microservices/order/internal/adapters/payment"
	"github.com/avi1152111680093/microservices/order/internal/application/core/api"
)

func main() {
	// dbAdapter, err := db.NewAdapter(config.GetDataSourceURL())
	dbAdapter, err := db.NewAdapter("root:verysecretpass@tcp(127.0.0.1:3306)/order")
	if err != nil{
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	paymentAdapter, err := payment.NewAdapter("localhost:3001")
	if err != nil {
		log.Fatalf("Failed to initialize payment stub. Error: %v", err)
	}

	fmt.Println("Connected to Payment Service..")

	application := api.NewApplication(dbAdapter, paymentAdapter)
	grpcAdapter := grpc.NewAdapter(application, 3000)
	grpcAdapter.Run()
}