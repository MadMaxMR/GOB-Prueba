package database

import (
	"log"

	"github.com/MadMaxMR/product-rest/models"

	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

func Connection() *gorm.DB {
	/*Coneccion con ElephantSQL*/
	//connStr := "postgres://arwpboxu:qP449bZjdC9jEpih47th8Hn21yi2Aj6h@motty.db.elephantsql.com/arwpboxu"
	/*Coneccion con Heroku*/
	connStr := "host=localhost port=5432 user=postgres dbname=productDB password=123456 sslmode=disable"
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
