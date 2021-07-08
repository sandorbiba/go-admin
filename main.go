package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sandorbiba/go-admin/database"
	"github.com/sandorbiba/go-admin/routes"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.Setup(app)

    app.Listen(":8000")
}