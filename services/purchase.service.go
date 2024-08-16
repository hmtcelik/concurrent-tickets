package services

import (
	"ticket-allocating/dal"
	"ticket-allocating/types"
	"ticket-allocating/utils"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CreatePurchase(c *fiber.Ctx) error {

	b := new(types.PurchaseCreate)

	if err := utils.ParseBodyAndValidate(c, b); err != nil {
		return err
	}

	d := &dal.Purchase{
		Quantity: b.Quantity,
		UserId:   b.UserId,
	}

	// For implementation, check the function in dal/purchase.dal.go
	if err := dal.CreatePurchaseWithTransaction(c.Params("id"), d); err != nil {
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
