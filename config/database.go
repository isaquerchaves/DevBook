package config

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDb() {
	var err error
	dsn := "host=motty.db.elephantsql.com user=cgisukrn password=8Q7NwgEyP32NIdmQiIpAVaJ6_M5aICem dbname=cgisukrn port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database")
	}
}
