package database

import (
	"log"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var AppDatabase *gorm.DB

func SetupDB() {
	db, err := gorm.Open(sqlite.Open("project_management.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database! \n", err)
		// return err
		os.Exit(2)
	}

	log.Println("Connected Successfully to Database")
	db.Logger = logger.Default.LogMode(logger.Info)
	log.Println("Running Migrations")

	db.AutoMigrate(&Project{}, &Task{}, &User{})

	log.Println("Project management application database created successfully.")

	AppDatabase = db

}
