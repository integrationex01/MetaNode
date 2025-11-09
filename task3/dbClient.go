package task3

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func DbClient() *gorm.DB {
	dsn := "host=localhost port=5432 user=user1 password=123456 dbname=test sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to the database successfully")
	return db
}
