package dal

import (
	"ticket-allocating/config/database"

	"gorm.io/gorm"
)

type Ticket struct {
	gorm.Model
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Allocation int    `json:"allocation"`
}

func CreateTicket(ticket *Ticket) *gorm.DB {
	return database.DB.Create(ticket)
}

func FindTicket(dest interface{}, conds ...interface{}) *gorm.DB {
	return database.DB.Model(&Ticket{}).Take(dest, conds...)
}
