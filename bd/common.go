package bd

import (
	"database/sql"
	"fmt"
	"os"

	"example.com/m/v2/models"
	"example.com/m/v2/secretm"
	_ "github.com/lib/pq"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

func ReadSecret() error {
	SecretModel, err = secretm.GetSecret(os.Getenv("SecretName"))
	return err
}

func DbConnect() error {
	Db, err = sql.Open("postgres", ConnStr(SecretModel))
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexión exitosa de la BD")
	return nil
}

func ConnStr(claves models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = claves.Username
	authToken = claves.Password
	dbEndpoint = claves.Host
	dbName = "postgres"

	// Cambia la cadena de conexión a PostgreSQL
	dsn := fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=require", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
