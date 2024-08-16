package routes

import (
	"ticket-allocating/services"

	"github.com/gofiber/fiber/v2"
)

func TicketRoutes(app fiber.Router) {
	r := app.Group("/tickets")

	r.Post("/", services.CreateTicket)
	r.Get("/:id", services.GetTicket)
	r.Post("/:id/purchases", services.CreatePurchase)

}
