package db

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DbInstance struct {
	Db *gorm.DB
}

var dbInstance *DbInstance

func connectToDb() {
	dsn := "breeze:breeze@tcp(127.0.0.1:3306)/breeze_nrc?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Error connecting to database!")
		os.Exit(2)
	}

	log.Println("Connected to database!")
	db.Logger = logger.Default.LogMode(logger.Info)

	// TODO: add migration
	log.Println("Migrating database...")

	dbInstance = &DbInstance{Db: db}
}
