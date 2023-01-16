package database

import (
	"errors"
)

func Delete(model interface{}, value string) error {

	db := GetConnection()
	defer db.Close()
	result := db.First(model, value)
	if result.RowsAffected == 0 {
		return errors.New("No se encontro datos con el ID: " + value)
	}
	err := db.Delete(model).Error
	if err != nil {
		return errors.New("Error al eliminar - " + err.Error())
	}
	return nil
}
