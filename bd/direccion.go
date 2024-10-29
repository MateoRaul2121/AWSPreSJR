package bd

import (
	"fmt"

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

	// Construcción de la sentencia SQL
	sentencia := "SELECT id_direcciones,id_secretaria, nombre, descripcion, activo, telefono, correo FROM areas.direcciones"
	var args []interface{}

	// Condiciones para la consulta
	if IDDirec > 0 {
		sentencia += " WHERE id_direcciones = $1"
		args = append(args, IDDirec)
	} else if len(Slug) > 0 {
		sentencia += " WHERE nombre ILIKE $1"
		args = append(args, "%"+Slug+"%")
	}

	fmt.Println("Consulta preparada:", sentencia)

	// Ejecución de la consulta
	rows, err := Db.Query(sentencia, args...)
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
