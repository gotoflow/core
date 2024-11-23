package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func GetPostgresConnection(host string, port string, username string, password string, database string) (*gorm.DB, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		host, port, username, database, password)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	return db, err
}

