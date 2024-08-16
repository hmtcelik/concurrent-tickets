package dal

import (
	"fmt"
	"ticket-allocating/config/database"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Purchase struct {
	gorm.Model
	Quantity int    `json:"quantity"`
	UserId   string `json:"user_id"`
}

// --------------- Logic ---------------
// Creates a purchase if enough allocation is available
// This function locks the ticket row for update to prevent concurrency issues
func CreatePurchaseWithTransaction(ticketID string, purchase *Purchase) error {
	return database.DB.Transaction(func(tx *gorm.DB) error {
		fmt.Println("CreatePurchaseWithTransaction")

		// Find the ticket and lock the row for update
		ticket := &Ticket{}
		if err := tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?", ticketID).First(ticket).Error; err != nil {
			fmt.Println("Ticket found")
			return gorm.ErrRecordNotFound
		}

		// Check if enough allocation is available
		if ticket.Allocation < purchase.Quantity {
			return gorm.ErrInvalidData
		}

		// Update the ticket allocation
		ticket.Allocation -= purchase.Quantity
		if err := tx.Save(ticket).Error; err != nil {
			return err
		}

		// Create the purchase
		if err := tx.Create(purchase).Error; err != nil {
			return err
		}

		fmt.Println("Purchase created")
		// After the transaction is committed, the ticket allocation will be updated
		// So after the return; the row will be unlocked because the transaction is committed
		return nil
	})
}
