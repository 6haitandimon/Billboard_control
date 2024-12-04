package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB
var DBADS *gorm.DB

func InitDB() {
	dsn := "root:root@tcp(localhost:5890)/MediaDB"
	dsnADS := "root:root@tcp(localhost:5890)/ADS"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	DB = db

	dbAds, err := gorm.Open(mysql.Open(dsnADS), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	DBADS = dbAds
	//err = DB.AutoMigrate(&models.User{})
	//if err != nil {
	//	log.Fatalf("Error migrating database: %v", err)
	//}
}
