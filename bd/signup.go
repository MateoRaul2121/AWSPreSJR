package bd

import (
	"fmt"

	"example.com/m/v2/models"
	_ "github.com/lib/pq" // Asegúrate de usar el driver para PostgreSQL
)

func SigUp(sig models.SignUp) error {
	fmt.Println("Comienza el registro")

	// Conectar a la base de datos
	err := DbConnect()
	if err != nil {
		return err
	}
	defer Db.Close()

	// Consulta con parámetros
	sentencia := `INSERT INTO usuarios.users ("USER_UUID", "USER_Email") VALUES ($1, $2)`
	fmt.Println(sentencia)

	// Ejecución de la consulta usando parámetros
	_, err = Db.Exec(sentencia, sig.UserUUID, sig.UserEmail)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Sig Up > Ejecución Exitosa")
	return nil
}
