package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DSN string = "host=localhost user=admin password=adminadmin dbname=tasks port=5432"
var DB *gorm.DB

func DBconnection() {
	var error error
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	if error != nil {
		log.Fatal(error)
		panic("Failed to connect to database")
	}

	log.Println("Database connected")

}
