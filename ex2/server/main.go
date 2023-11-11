package main

import (
	"go_admin/config"
	"go_admin/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	server := ":8002"
	db := config.NewDatabaseMariaDB()
	// config.GenDatSimple(db)
	app := fiber.New()
	routes.AdminRouter(app, db)
	app.Listen(server)

}
