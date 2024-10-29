package routers

import (
	"encoding/json"
	"strconv"

	"example.com/m/v2/bd"
	"example.com/m/v2/models"
	"github.com/aws/aws-lambda-go/events"
)

func InsertJefatura(body string, User string) (int, string) {
	var t models.Jefatura

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if t.DirecID == 0 {
		return 400, "Debe especificar el ID de la direccion"
	}
	if len(t.JefaNombre) == 0 {
		return 400, "Debe especificar el nombre de la Jefatura"
	}
	if len(t.JefaDescripcion) == 0 {
		return 400, "Debe especificar la descripcion de la Jefatura"
	}
	if len(t.JefaTelefono) == 0 {
		return 400, "Debe especificar el telefono de la Jefatura"
	}
	if len(t.JefaCorreo) == 0 {
		return 400, "Debe especificar el correo de la Jefatura"
	}

	result, err := bd.InsertJefatura(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la jefatura " + t.JefaNombre + " > " + err.Error()
	}
	return 200, "{ JefatutaID: " + result + "}"

}

func SelectJefatura(body string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var id_jefa int
	var Slug string

	if len(request.QueryStringParameters["id_jefatura"]) > 0 {
		id_jefa, err = strconv.Atoi(request.QueryStringParameters["id_jefatura"])
		if err != nil {
			return 500, "Ocurrio un error al intentar convertir en entero el valor " + request.QueryStringParameters["id_jefatura"]
		}
	} else {
		if len(request.QueryStringParameters["slug"]) > 0 {
			Slug = request.QueryStringParameters["slug"]
		}
	}

	lista, err2 := bd.SelectJefatura(id_jefa, Slug)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar capturar Jefatura/s > " + err2.Error()
	}

	Jefa, err3 := json.Marshal(lista)
	if err3 != nil {
		return 400, "Ocurrio un error al intentar convertir en JSON Jefatura/s > " + err3.Error()
	}

	return 200, string(Jefa)
}
