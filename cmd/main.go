package main

import (
	"log"
	"os"

	"go-hex/internal/adapters/app"
	"go-hex/internal/adapters/core/arithmetic"
	"go-hex/internal/adapters/framework/left/grpc"
	"go-hex/internal/adapters/framework/right/db"
	"go-hex/internal/ports"
)

func main() {
	var core ports.ArithmeticPort = arithmetic.NewAdapter()

	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		log.Fatal("DB_DSN environment variable is required")
	}

	dbAdapter, err := db.NewAdapter("mysql", dbDSN)
	if err != nil {
		panic(err)
	}

	var api ports.APIPort = api.NewAdapter(core, dbAdapter)

	grpcAdapter := grpc.NewAdapter(api)
	grpcAdapter.Run()
}
