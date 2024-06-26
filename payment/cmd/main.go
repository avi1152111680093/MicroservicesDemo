package main

import (
	"log"

	"github.com/avi1152111680093/microservices/payment/internal/adapters/db"
	"github.com/avi1152111680093/microservices/payment/internal/adapters/grpc"
	"github.com/avi1152111680093/microservices/payment/internal/application/core/api"
)

func main () {
	dbAdapter, err := db.NewAdapter("root:verysecretpass@tcp(127.0.0.1:3306)/payment")
	if err != nil {
		log.Fatalf("Failed to connect to database. Error: %v", err)
	}

	application := api.NewApplication(dbAdapter)
	grpcAdapter := grpc.NewAdapter(application, 3001)
	grpcAdapter.Run()
}