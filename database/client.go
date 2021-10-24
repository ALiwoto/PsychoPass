package database

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

var SESSION *gorm.DB

func StartDatabase() {
	db, err := gorm.Open(sqlite.Open(fmt.Sprintf("sibyl.db")), &gorm.Config{})
	if err != nil {
		log.Fatalln("failed to connect database")
	}

	SESSION = db
	log.Println("Database connected")

	// Create tables if they don't exist
	err = SESSION.AutoMigrate(&User{}, &Token{})
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Auto-migrated database schema")

}
