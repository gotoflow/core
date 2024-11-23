package database

import (
	"log"

	"gorm.io/gorm"
)

func GetDatabase(driver string, host string, port string, username string, password string, database string) *gorm.DB {
	switch driver {
		case "postgres":
			db, err := GetPostgresConnection(host, port, username, password, database)
			if err != nil {
				log.Fatal(err)
			}
			return db
		default:
			log.Fatal("Invalid driver")
			return nil
	}
}