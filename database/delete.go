package database

import (
	"errors"
)

func Delete(model interface{}, id string) (string, error) {

	db := GetConnection()
	defer db.Close()
	result := db.Find(model, id)
	if result.RowsAffected == 0 {
		return "", errors.New("No se encontro datos con el ID: " + id)
	}
	err := db.Delete(model, id).Error
	if err != nil {
		return "", errors.New("Error al eliminar - " + err.Error())
	}
	return "Elemento eliminado correctamente", nil
}
