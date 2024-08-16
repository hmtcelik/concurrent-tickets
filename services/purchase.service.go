package services

import (
	"strconv"
	"ticket-allocating/dal"
	"ticket-allocating/types"
	"ticket-allocating/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// Create Purchase
//
//	@Summary		Create Purchase
//	@Description	Creating a new Purchase
//	@Tags			Purchases
//	@Accept			json
//	@Produce		json
//	@Param			id		path	string					true	"Ticket ID"
//	@Param			body	body	types.PurchaseCreate	true	"Purchase Create"
//	@Success		201
//	@Router			/tickets/{id}/purchases [post]
func CreatePurchase(c *fiber.Ctx) error {

	b := new(types.PurchaseCreate)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	d := &dal.Purchase{
		Quantity: b.Quantity,
		UserId:   b.UserId,
	}

	ticktedId, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Please provide a valid ticket id")
	}

	// For implementation, check the function in dal/purchase.dal.go
	if err := dal.CreatePurchaseWithTransaction(ticktedId, d); err != nil {
		if err == gorm.ErrRecordNotFound {
			return fiber.NewError(fiber.StatusNotFound, err.Error())
		}
		if err == gorm.ErrInvalidData {
			return fiber.NewError(fiber.StatusBadRequest, "Not enough tickets available")
		}
		return fiber.NewError(fiber.StatusInternalServerError, err.Error())
	}

	return c.Status(fiber.StatusCreated).Send(nil)
}
