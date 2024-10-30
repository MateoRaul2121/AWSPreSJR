package bd

import (
	"fmt"
	"strconv"

	"example.com/m/v2/models"
)

func InsertDireccion(c models.Direccion) (string, error) {
	fmt.Println("Comienza Registro de Direccion")
	// Conectar a la base de datos
	err := DbConnect()
	if err != nil {
		return "", err
	}
	defer Db.Close()
	// Sentencia para llamar al procedimiento almacenado
	sentencia := `SELECT * FROM areas.create_direc($1,$2,$3,$4,$5,$6)`
	DireccionRetornada := ""
	err = Db.QueryRow(sentencia, c.SecreID, c.DirecNombre, c.DirecDescripcion, c.DirecActivo, c.DirecTelefono, c.DirecCorreo).Scan(&DireccionRetornada)
	if err != nil {
		return "", err
	}

	fmt.Println("Insert Direccion > Ejecución Exitosa")
	return DireccionRetornada, nil
}

func SelectDireccion(IDDirec int, Slug string) ([]models.Direccion, error) {
	fmt.Println("Comienza Select Direccción")

	var Direc []models.Direccion

	err := DbConnect()
	if err != nil {
		return Direc, err
	}
	defer Db.Close()

	// Ajustar los parámetros para la llamada a la función almacenada
	var idParam interface{}
	var slugParam interface{}

	if IDDirec > 0 {
		idParam = IDDirec
	} else {
		idParam = nil
	}

	if len(Slug) > 0 {
		slugParam = Slug
	} else {
		slugParam = nil
	}

	// Construcción de la llamada a la función almacenada
	sentencia := "SELECT * FROM areas.get_direc($1, $2)"
	rows, err := Db.Query(sentencia, idParam, slugParam)
	if err != nil {
		return Direc, err
	}
	defer rows.Close()

	// Lectura de resultados
	for rows.Next() {
		var b models.Direccion
		err := rows.Scan(&b.DirecID, &b.SecreID, &b.DirecNombre, &b.DirecDescripcion, &b.DirecActivo, &b.DirecTelefono, &b.DirecCorreo)
		if err != nil {
			return Direc, err
		}
		Direc = append(Direc, b)
	}

	// Comprobación de errores en la iteración
	if err := rows.Err(); err != nil {
		return Direc, err
	}

	fmt.Println("Select Direccion > Ejecución Exitosa")
	return Direc, nil
}

func DeleteFisicoDireccion(id int) error {
	fmt.Println("Comienza Registro de Delete Direccion")

	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	sentencia := "DELETE FROM areas.direcciones WHERE id_direcciones = " + strconv.Itoa(id)

	_, err = Db.Exec(sentencia)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println(sentencia)
	fmt.Println("Delete Direccion > Ejecución Exitosa")
	return nil
}

func UpdateDireccion(c models.Direccion) error {
	fmt.Println("Comienza Registro de Update Direccion")

	// Conexión a la base de datos
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	// Llamada a la función almacenada para actualizar los datos de Direcciones
	sentencia := "SELECT areas.update_direccion($1, $2, $3, $4, $5, $6, $7)"
	_, err = Db.Exec(sentencia, c.DirecID, c.SecreID, c.DirecNombre, c.DirecDescripcion, c.DirecActivo, c.DirecTelefono, c.DirecCorreo)
	if err != nil {
		fmt.Println("Error al ejecutar la actualización:", err.Error())
		return err
	}

	fmt.Println("Update Direccion > Ejecución Exitosa")
	return nil
}
