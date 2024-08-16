package main

import (
	"ticket-allocating/config/database"
	"ticket-allocating/dal"
	"ticket-allocating/routes"
	"ticket-allocating/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "ticket-allocating/docs"
)

// @title			Ticket Allocating API
// @version		1.0
// @description	Concurrent Ticket Allocating API with Golang Fiber
// @contact.name	API Support
// @contact.email	abhmtcelik@gmail.com
// @host			localhost:3000
// @BasePath		/
func main() {
	database.Connect()
	database.Migrate(&dal.Ticket{}, &dal.Purchase{})

	app := fiber.New(fiber.Config{
		ErrorHandler: utils.ErrorHandler,
	})

	// app routes:
	routes.TicketRoutes(app)

	// swagger
	app.Get("/swagger/*", swagger.HandlerDefault)

	app.Listen(":3000")
}
