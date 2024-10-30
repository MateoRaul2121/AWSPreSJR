package bd

import (
	"fmt"

	"example.com/m/v2/models"
)

func InsertSecretaria(c models.Secretaria) (string, error) {
	fmt.Println("Comienza Registro de Secretarias")
	// Conectar a la base de datos
	err := DbConnect()
	if err != nil {
		return "", err
	}
	defer Db.Close()
	// Sentencia para llamar al procedimiento almacenado
	sentencia := `SELECT * FROM areas.create_sec($1,$2,$3,$4,$5)`
	secretariaRetornada := ""
	err = Db.QueryRow(sentencia, c.SecreNombre, c.SecreDescripcion, c.SecreActivo, c.SecreTelefono, c.SecreCorreo).Scan(&secretariaRetornada)
	if err != nil {
		return "", err
	}

	fmt.Println("Insert Secretaria > Ejecución Exitosa")
	return secretariaRetornada, nil
}

func SelectSecretaria(IDSecre int, Slug string) ([]models.Secretaria, error) {
	fmt.Println("Comienza Select Secretaria")

	var Secr []models.Secretaria

	err := DbConnect()
	if err != nil {
		return Secr, err
	}
	defer Db.Close()

	// Ajustar los parámetros para la llamada a la función almacenada
	var idParam interface{}
	var slugParam interface{}

	if IDSecre > 0 {
		idParam = IDSecre
	} else {
		idParam = nil
	}

	if len(Slug) > 0 {
		slugParam = Slug
	} else {
		slugParam = nil
	}

	// Construcción de la llamada a la función almacenada
	sentencia := "SELECT * FROM areas.get_secretarias($1, $2)"
	rows, err := Db.Query(sentencia, idParam, slugParam)
	if err != nil {
		return Secr, err
	}
	defer rows.Close()

	// Lectura de resultados
	for rows.Next() {
		var b models.Secretaria
		err := rows.Scan(&b.SecreID, &b.SecreNombre, &b.SecreDescripcion, &b.SecreActivo, &b.SecreTelefono, &b.SecreCorreo)
		if err != nil {
			return Secr, err
		}
		Secr = append(Secr, b)
	}

	// Comprobación de errores en la iteración
	if err := rows.Err(); err != nil {
		return Secr, err
	}

	fmt.Println("Select Secretaria > Ejecución Exitosa")
	return Secr, nil
}
