package main

import (
	"backend-umkm/database"
	"backend-umkm/route"

	"github.com/gofiber/fiber/v2"
)

func main() {
	//initiate db
	database.DatabaseInit()
	app := fiber.New()

	route.RouteInit(app)

	app.Listen(":3000")
}
