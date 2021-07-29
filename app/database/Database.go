package database

import (
	"log"

	"github.com/Artpou/elastic-search-go/app/models"
	_ "github.com/elastic/go-elasticsearch/v7"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB
var err error
var dbServer, dbName, dbUsername, dbPassword, dbPort, dbConn string

func InitDb() *gorm.DB {
	dbServer = "localhost"
	dbName = "localhost"
	dbPort = "8080"
	dbConn = "@tcp(" + dbServer + ":" + dbPort + ")/" + dbName + "?charset=utf8&parseTime=True"

	db, err = gorm.Open("mysql", dbConn)

	if err != nil {
		log.Println("DB connection Failed to Open")
	} else {
		log.Println("DB connection Established")
		db.AutoMigrate(&models.User{})
	}
	return db
}

func GetDb() *gorm.DB {
	return db
}
