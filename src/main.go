package main

import (
	"ecommerce/user/usermanagement/src/config"
	"ecommerce/user/usermanagement/src/db"
	"ecommerce/user/usermanagement/src/server"
	"log"
)

func main() {
	ctx := config.StartupContext()
	if err := db.LoadConnection(); err != nil {
		log.Fatalf("Failed to connect to DB %v", err)
		return
	}
	server.Start(ctx)
}
