package bd

import (
	"fmt"
	"strconv"

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

func DeleteFisicoJefatura(id int) error {
	fmt.Println("Comienza Registro de Delete Jefatura")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "DELETE FROM areas.jefaturas WHERE id_jefatura = " + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(sentencia)
	fmt.Println("Delete Jefatura > Ejecución Exitosa")
	return nil
}

func UpdateJefatura(c models.Jefatura) error {
	fmt.Println("Comienza Registro de Update Jefatura")

	// Conexión a la base de datos
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	// Llamada a la función almacenada para actualizar los datos de Jefaturas
	sentencia := "SELECT areas.update_jefatura($1, $2, $3, $4, $5, $6, $7)"
	_, err = Db.Exec(sentencia, c.JefaID, c.DirecID, c.JefaNombre, c.JefaDescripcion, c.JefaActivo, c.JefaTelefono, c.JefaCorreo)
	if err != nil {
		fmt.Println("Error al ejecutar la actualización:", err.Error())
		return err
	}

	fmt.Println("Update Jefatura > Ejecución Exitosa")
	return nil
}
