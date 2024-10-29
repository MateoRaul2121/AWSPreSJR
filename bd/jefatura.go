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

	err := DbConnect()
	if err != nil {
		return Jefa, err
	}
	defer Db.Close()

	// Construcción de la sentencia SQL
	sentencia := "SELECT id_jefatura,id_direcciones, nombre, descripcion, activo, telefono, correo FROM areas.jefaturas"
	var args []interface{}

	// Condiciones para la consulta
	if IDJefa > 0 {
		sentencia += " WHERE id_jefatura = $1"
		args = append(args, IDJefa)
	} else if len(Slug) > 0 {
		sentencia += " WHERE nombre ILIKE $1"
		args = append(args, "%"+Slug+"%")
	}

	fmt.Println("Consulta preparada:", sentencia)

	// Ejecución de la consulta
	rows, err := Db.Query(sentencia, args...)
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
