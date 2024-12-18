package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var DB *gorm.DB
var DBADS *gorm.DB

func InitDB() {
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	mediaDB := os.Getenv("MEDIA_DB")
	adsDB := os.Getenv("ADS_DB")

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, mediaDB)

	dsnADS := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, adsDB)
	//dsn := "root:root@tcp(localhost:5890)/MediaDB"
	//dsnADS := "root:root@tcp(localhost:5890)/ADS"

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
