package database

import (
	"errors"
	//"fmt"
	"strconv"
)

func GetAll(model interface{}, page string) error {
	db := GetConnection()
	defer db.Close()
	pageInt, _ := strconv.Atoi(page)
	if page == "" {
		db.Set("gorm:order_by_primary_key", "ASC").Find(model)
		return nil
	}
	if page == "1" {
		result := db.Limit(10).Set("gorm:order_by_primary_key", "ASC").Find(model)
		if result.RowsAffected == 0 {
			return errors.New("No se encontro datos en la página: " + page)
		}
		return nil

	} else {
		result := db.Limit(10).Offset((pageInt-1)*10).Set("gorm:order_by_primary_key", "ASC").Find(model)
		if result.RowsAffected == 0 {
			return errors.New("No se encontro datos en la página: " + page)
		}
		return nil
	}
}

func Get(model interface{}, value string) error {
	db := GetConnection()
	defer db.Close()

	result := db.Debug().First(model, value)
	if result.RowsAffected == 0 {
		return errors.New("No se encontro datos con el ID: " + value)
	}
	return nil
}
