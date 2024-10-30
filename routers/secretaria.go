package routers

import (
	"encoding/json"
	"strconv"

	"example.com/m/v2/bd"
	"example.com/m/v2/models"
	"github.com/aws/aws-lambda-go/events"
)

func InsertSecretaria(body string, User string) (int, string) {
	var t models.Secretaria

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.SecreNombre) == 0 {
		return 400, "Debe especificar el nopmbre de la secreataria"
	}
	if len(t.SecreDescripcion) == 0 {
		return 400, "Debe especificar la descripcion de la secreataria"
	}
	if len(t.SecreTelefono) == 0 {
		return 400, "Debe especificar el telefono de la secretaria"
	}
	if len(t.SecreCorreo) == 0 {
		return 400, "Debe especificar el correo de la secreataria"
	}

	result, err := bd.InsertSecretaria(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la secretaria " + t.SecreNombre + " > " + err.Error()
	}
	return 200, "{ SecretariaNombre: " + result + "}"

}

func SelectSecretaria(body string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var id_secre int
	var Slug string

	if len(request.QueryStringParameters["id_secretaria"]) > 0 {
		id_secre, err = strconv.Atoi(request.QueryStringParameters["id_secretaria"])
		if err != nil {
			return 500, "Ocurrio un error al intentar convertir en entero el valor " + request.QueryStringParameters["id_secretaria"]
		}
	} else {
		if len(request.QueryStringParameters["slug"]) > 0 {
			Slug = request.QueryStringParameters["slug"]
		}
	}

	lista, err2 := bd.SelectSecretaria(id_secre, Slug)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar capturar Secretaria/s > " + err2.Error()
	}

	Secr, err3 := json.Marshal(lista)
	if err3 != nil {
		return 400, "Ocurrio un error al intentar convertir en JSON Secretaria/s > " + err3.Error()
	}

	return 200, string(Secr)
}

func DeleteFisicoSecretaria(body string, User string, id int) (int, string) {
	if id == 0 {
		return 400, "Debe especificar ID de la Secretaria a Borrar"
	}

	err := bd.DeleteFisicoSecretaria(id)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el DELETE de la secretaria " + strconv.Itoa(id) + " > " + err.Error()
	}
	return 200, "Delete OK"
}

func UpdateSecretaria(body string, User string, id int) (int, string) {
	var t models.Secretaria

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	t.SecreID = id
	err2 := bd.UpdateSecretaria(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el UPDATE de la Secretaria " + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Update OK"
}
