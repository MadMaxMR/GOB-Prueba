package database

import (
	"errors"
)

func Create(model interface{}) error {
	db := GetConnection()
	defer db.Close()
	err := db.Create(model).Error
	if err != nil {
		return errors.New("Error al guardar - " + err.Error())
	}
	return nil
}
