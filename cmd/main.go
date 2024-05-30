package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/ilhamgepe/prakerja-s7/cmd/api"
	"github.com/ilhamgepe/prakerja-s7/config"
	"github.com/ilhamgepe/prakerja-s7/db/postgres"
)

func init() {
	if err := config.LoadConfig("./.env"); err != nil {
		log.Fatalf("failed to load config: %v", err)
	}
}

func main() {
	db := postgres.NewPostgresDB()
	server := api.NewServer(config.Get.Addr, db)

	go func() {
		log.Fatal(server.Run())
	}()

	fmt.Println("server running on port http://localhost:" + config.Get.Addr)

	// Graceful Shutdown
	stopC := make(chan os.Signal, 1)
	signal.Notify(stopC, os.Interrupt)
	fmt.Println("signal received: ", <-stopC)

	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatalf("server shutdown error: %s", err)
	}

	fmt.Println("server shutdown successfully")
	os.Exit(0)
}
