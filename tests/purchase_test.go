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

	// Mocking the ticket row
	rows := sqlmock.NewRows([]string{"id", "name", "desc", "allocation"}).
		AddRow(1, "Test Ticket", "Test Description", 100)

	// Expectations
	mock.ExpectBegin()
	mock.ExpectQuery(`SELECT \* FROM "tickets" WHERE id = \$1 AND "tickets"\."deleted_at" IS NULL ORDER BY "tickets"\."id" LIMIT \$2 FOR UPDATE`).
		WithArgs(1, 1).
		WillReturnRows(rows)

	newAllocation := 100 - 2
	mock.ExpectExec(`UPDATE "tickets" SET "created_at"=\$1,"updated_at"=\$2,"deleted_at"=\$3,"name"=\$4,"desc"=\$5,"allocation"=\$6 WHERE "tickets"."deleted_at" IS NULL AND "id" = \$7`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, "Test Ticket", "Test Description", newAllocation, 1).
		WillReturnResult(sqlmock.NewResult(1, 1))

	mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO "purchases" ("created_at","updated_at","deleted_at","quantity","user_id") VALUES ($1,$2,$3,$4,$5) RETURNING "id"`)).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), nil, 2, "user1").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	// Running the create purchase function
	purchase := &dal.Purchase{
		Quantity: 2,
		UserId:   "user1",
	}
	err := dal.CreatePurchaseWithTransaction(1, purchase)
	assert.NoError(t, err, "Failed to create purchase")

	// Verify the expectations
	err = mock.ExpectationsWereMet()
	assert.NoError(t, err, "Expectations were not met")
}
