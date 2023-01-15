package database

import (
	"log"

	"github.com/MadMaxMR/Products-Restful/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var host = "localhost"
var port = "5432"
var user = "postgres"
var dbname = "productDB"
var password = "123456"

func Connection() *gorm.DB {

	connStr := host + " " + port + " " + user + " " + dbname + " " + password + " sslmode=disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

//utilizando el patron de diseño Singleton
func GetConnection() *gorm.DB {
	db := Connection()
	return db
}

func Migrate() {
	db := GetConnection()
	defer db.Close()

	log.Printf("Migrando base de datos")

	db.AutoMigrate(&models.Product{}, &models.OtherImages{})
}
