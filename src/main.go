package main

import (
	"ecommerce/user/usermanagement/src/config"
	"ecommerce/user/usermanagement/src/server"
)

func main() {
	ctx := config.StartupContext()
	server.Start(ctx)
}
