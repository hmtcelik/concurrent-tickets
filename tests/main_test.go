package tests

import (
	"testing"
	"ticket-allocating/config/database"
	"ticket-allocating/dal"
)

func TestMain(m *testing.M) {
	// Connect to an in-memory database for testing
	database.Connect("file::memory:?cache=shared&_timeout=10000")
	database.Migrate(&dal.Ticket{}, &dal.Purchase{})

	m.Run()
}
