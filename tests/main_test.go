package tests

import (
	"database/sql"
	"testing"
	"ticket-allocating/config/database"

	"github.com/DATA-DOG/go-sqlmock"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// This function is special, its like a constructor for the tests
// You can do whatever you want before the tests (or after)
func TestMain(m *testing.M) {
	m.Run()
}

// This function is used to connect to a mock database for testing
// So that the actual database is not affected by the tests
func ConnectMockDB(t *testing.T) (*sql.DB, sqlmock.Sqlmock) {
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	dialector := postgres.New(postgres.Config{
		Conn:       mockDb,
		DriverName: "postgres",
	})
	db, err := gorm.Open(dialector, &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a gorm database connection", err)
	}

	database.DB = db

	return mockDb, mock
}
