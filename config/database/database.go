package database

import (
	"fmt"
	"ticket-allocating/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Connect() {
	// Data Source Name
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_PASS, config.DB_NAME, config.DB_PORT,
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger:  logger.Default.LogMode(logger.Error),
		NowFunc: func() time.Time { return time.Now().Local() },
	})

	if err != nil {
		fmt.Println("[DATABASE]::CONNECTION_ERROR")
		panic(err)
	}

	// Setting the database connection to use in routes
	DB = db
	fmt.Println("[DATABASE]::CONNECTED")
}

// Migrate migrates all the database tables
func Migrate(tables ...interface{}) error {
	return DB.AutoMigrate(tables...)
}
