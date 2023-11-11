package main

import (
	"go_beer/config"
	"go_beer/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := ":8001"
	mariaDB := config.NewDatabaseMariaDB()
	mongoDB := config.NewDatabaseMongoDB()
	app := fiber.New()
	routes.BeerRouter(app, mariaDB, mongoDB)
	app.Listen(server)
}
