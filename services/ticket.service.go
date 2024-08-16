package services

import (
	"ticket-allocating/dal"
	"ticket-allocating/types"
	"ticket-allocating/utils"

	"github.com/gofiber/fiber/v2"
)

func CreateTicket(c *fiber.Ctx) error {

	b := new(types.TicketCreate)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	d := &dal.Ticket{
		Name:       b.Name,
		Desc:       b.Desc,
		Allocation: *b.Allocation,
	}

	if err := dal.CreateTicket(d).Error; err != nil {
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).JSON(&types.TicketResponse{
		ID:         d.ID,
		Name:       d.Name,
		Desc:       d.Desc,
		Allocation: d.Allocation,
	})
}

func GetTicket(c *fiber.Ctx) error {

	id := c.Params("id")

	d := new(dal.Ticket)

	if err := dal.FindTicket(d, id).Error; err != nil {
		return fiber.NewError(fiber.StatusNotFound, err.Error())
	}

	return c.JSON(&types.TicketResponse{
		ID:         d.ID,
		Name:       d.Name,
		Desc:       d.Desc,
		Allocation: d.Allocation,
	})
}
