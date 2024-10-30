package bd

import (
	"fmt"

	"example.com/m/v2/models"
)

func InsertJefatura(c models.Jefatura) (string, error) {
	fmt.Println("Comienza Registro de Jefatura")
	// Conectar a la base de datos
	err := DbConnect()
	if err != nil {
		return "", err
	}
	defer Db.Close()
	// Sentencia para llamar al procedimiento almacenado
	sentencia := `SELECT * FROM areas.create_jefat($1,$2,$3,$4,$5,$6)`
	JefaturaRetornada := ""
	err = Db.QueryRow(sentencia, c.DirecID, c.JefaNombre, c.JefaDescripcion, c.JefaActivo, c.JefaTelefono, c.JefaCorreo).Scan(&JefaturaRetornada)
	if err != nil {
		return "", err
	}

	fmt.Println("Insert Jefatura > Ejecución Exitosa")
	return JefaturaRetornada, nil
}

func SelectJefatura(IDJefa int, Slug string) ([]models.Jefatura, error) {
	fmt.Println("Comienza Select Jefatura")

	var Jefa []models.Jefatura

	// Conexión a la base de datos
	err := DbConnect()
	if err != nil {
		return Jefa, err
	}
	defer Db.Close()

	// Ajustar los parámetros para la llamada a la función almacenada
	var idParam interface{}
	var slugParam interface{}

	if IDJefa > 0 {
		idParam = IDJefa
	} else {
		idParam = nil
	}

	if len(Slug) > 0 {
		slugParam = Slug
	} else {
		slugParam = nil
	}

	// Construcción de la llamada a la función almacenada
	sentencia := "SELECT * FROM areas.get_jefatura($1, $2)"
	rows, err := Db.Query(sentencia, idParam, slugParam)
	if err != nil {
		return Jefa, err
	}
	defer rows.Close()

	// Lectura de resultados
	for rows.Next() {
		var b models.Jefatura
		err := rows.Scan(&b.JefaID, &b.DirecID, &b.JefaNombre, &b.JefaDescripcion, &b.JefaActivo, &b.JefaTelefono, &b.JefaCorreo)
		if err != nil {
			return Jefa, err
		}
		Jefa = append(Jefa, b)
	}

	// Comprobación de errores en la iteración
	if err := rows.Err(); err != nil {
		return Jefa, err
	}

	fmt.Println("Select Jefatura > Ejecución Exitosa")
	return Jefa, nil
}
