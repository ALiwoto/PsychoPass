package database

import (
	"fmt"

	"gitlab.com/Dank-del/SibylAPI-Go/core/sibylConfig"
	"gitlab.com/Dank-del/SibylAPI-Go/core/utils/logging"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var SESSION *gorm.DB

func StartDatabase() {
	var db *gorm.DB
	var err error
	if sibylConfig.SibylConfig.UseSqllite {
		db, err = gorm.Open(sqlite.Open(
			fmt.Sprintf("%s.db", sibylConfig.SibylConfig.DbName)), &gorm.Config{})
	} else {
		db, err = gorm.Open(postgres.Open(sibylConfig.SibylConfig.DbUrl), &gorm.Config{})
	}

	if err != nil {
		logging.Fatal("failed to connect to the database:", err)
	}

	SESSION = db
	logging.Info("Database connected")

	// Create tables if they don't exist
	err = SESSION.AutoMigrate(&User{}, &Token{})
	if err != nil {
		logging.Fatal(err)
	}

	logging.Info("Auto-migrated database schema")
}
