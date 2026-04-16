package main

import (
	"log"
	"os"

	api "go-hex/internal/adapters/app"
	"go-hex/internal/adapters/core/arithmetic"
	"go-hex/internal/adapters/framework/left/grpc"
	"go-hex/internal/adapters/framework/right/db"
	"go-hex/internal/ports"
)

func main() {
	dbDriver := os.Getenv("DB_DRIVER")
	dbDSN := os.Getenv("DB_DSN")
	if dbDSN == "" {
		log.Fatal("DB_DSN environment variable is required")
	}

	dbAdapter, err := db.NewAdapter(dbDriver, dbDSN)
	if err != nil {
		panic(err)
	}
	defer dbAdapter.CloseDbConnection()

	var arithAdapter ports.ArithmeticPort = arithmetic.NewAdapter()
	var api ports.APIPort = api.NewAdapter(arithAdapter, dbAdapter)

	grpcAdapter := grpc.NewAdapter(api)
	grpcAdapter.Run()
}
