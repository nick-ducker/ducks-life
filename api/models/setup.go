package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var DB *gorm.DB
var dbName string

func ConnectDatabase(testing bool) *gorm.DB {
	dbName := "prod.db"
	if testing {
		dbName = "test.db"
	}

	database, err := gorm.Open("sqlite3", dbName)

	if err != nil {
		panic("Failed to connect to database!")
	} else {
		fmt.Printf("DB connection established")
	}

	database.AutoMigrate(&Rambling{})

	DB = database
	return DB
}

func DeleteDB(database *gorm.DB) error {
	database.Close()
	err := os.Remove(dbName)
	return err
}
