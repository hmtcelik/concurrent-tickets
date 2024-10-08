package tests

import (
	"regexp"
	"testing"
	"ticket-allocating/dal"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestGetTicket(t *testing.T) {
	db, mock := ConnectMockDB(t)
	defer db.Close()

	// Mocking the ticket row
	rows := sqlmock.NewRows([]string{"id", "name", "desc", "allocation"}).
		AddRow(1, "Test Ticket", "Test Description", 5)

	// Expectations
	mock.ExpectQuery("^SELECT (.+) FROM \"tickets\"*").
		WillReturnRows(rows)

	// Running the find ticket function
	ticket := &dal.Ticket{}
	err := dal.FindTicket(ticket, 1).Error
	assert.NoError(t, err, "Failed to find ticket")

	// Verify the expectations
	assert.Equal(t, "Test Ticket", ticket.Name, "Name mismatch")
	assert.Equal(t, "Test Description", ticket.Desc, "Description mismatch")
	assert.Equal(t, 5, ticket.Allocation, "Allocation mismatch")

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "Expectations were not met")
}

func TestCreateTicket(t *testing.T) {
	db, mock := ConnectMockDB(t)
	defer db.Close()

	ticket := &dal.Ticket{
		Name:       "Test Ticket",
		Desc:       "Test Description",
		Allocation: 100,
	}

	// Expectations
	mock.ExpectBegin()
	mock.ExpectQuery(regexp.QuoteMeta("INSERT INTO \"tickets\" (\"created_at\",\"updated_at\",\"deleted_at\",\"name\",\"desc\",\"allocation\") VALUES ($1,$2,$3,$4,$5,$6) RETURNING \"id\"")).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, ticket.Name, ticket.Desc, ticket.Allocation).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(ticket.ID))
	mock.ExpectCommit()

	// Running the create ticket function
	err := dal.CreateTicket(ticket).Error
	assert.NoError(t, err, "Failed to create ticket")

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "Expectations were not met")
}
