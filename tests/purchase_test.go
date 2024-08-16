package tests

import (
	"regexp"
	"testing"
	"ticket-allocating/dal"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestPurchaseTicket(t *testing.T) {
	db, mock := ConnectMockDB(t)
	defer db.Close()

	// Mock finding the ticket
	rows := sqlmock.NewRows([]string{"id", "name", "desc", "allocation"}).
		AddRow(1, "Test Ticket", "Test Description", 100)

	// Ensure the query matches exactly
	mock.ExpectQuery("^SELECT \\* FROM \"tickets\" WHERE \"tickets\".\"id\" = \\$1 AND \"tickets\".\"deleted_at\" IS NULL LIMIT \\$2$").
		WithArgs(1, 1).
		WillReturnRows(rows)

	ticket := &dal.Ticket{}
	err := dal.FindTicket(ticket, 1).Error
	assert.NoError(t, err, "Failed to find ticket")

	purchase := &dal.Purchase{
		Quantity: 2,
		UserId:   "user1",
	}

	mock.ExpectBegin()

	// Mock locking the ticket row
	mock.ExpectQuery("^SELECT \\* FROM \"tickets\" WHERE id = \\$1 AND \"tickets\".\"deleted_at\" IS NULL ORDER BY \"tickets\".\"id\" LIMIT \\$2 FOR UPDATE$").
		WithArgs(1, 1).
		WillReturnRows(rows)

	// Mock inserting the purchase
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO \"purchases\" (\"created_at\",\"updated_at\",\"deleted_at\",\"quantity\",\"user_id\") VALUES ($1,$2,$3,$4,$5) RETURNING \"id\"")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, purchase.Quantity, purchase.UserId).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// Mock updating the ticket's allocation
	newAllocation := ticket.Allocation - purchase.Quantity
	mock.ExpectExec(regexp.QuoteMeta("UPDATE \"tickets\" SET \"allocation\" = $1 WHERE \"id\" = $2")).
		WithArgs(newAllocation, ticket.ID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Mock the end of the transaction
	mock.ExpectCommit()

	// Perform the actual test
	err = dal.CreatePurchaseWithTransaction(int(ticket.ID), purchase)
	assert.NoError(t, err, "Failed to create purchase")

	// Verify all expectations were met
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "Expectations were not met")
}
