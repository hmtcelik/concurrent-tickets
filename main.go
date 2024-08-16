package main

import (
	"ticket-allocating/config/database"
	"ticket-allocating/dal"
	"ticket-allocating/routes"
	"ticket-allocating/utils"

	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()
	database.Migrate(&dal.Ticket{}, &dal.Purchase{})

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	routes.TicketRoutes(app)

	app.Listen(":3000")
}
