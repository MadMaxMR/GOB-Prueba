package database

import (
	"errors"
)

func Update(model interface{}, value string) error {
	db := GetConnection()
	defer db.Close()
	err := db.Model(model).Where(value).Update(model).Error

	if err != nil {
		return errors.New("Error al actualizar - " + err.Error())
	}
	return nil

}
