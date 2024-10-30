package routers

import (
	"encoding/json"
	"strconv"

	"example.com/m/v2/bd"
	"example.com/m/v2/models"
	"github.com/aws/aws-lambda-go/events"
)

func InsertDireccion(body string, User string) (int, string) {
	var t models.Direccion

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(t.DirecNombre) == 0 {
		return 400, "Debe especificar el nombre de la direccion"
	}
	if len(t.DirecDescripcion) == 0 {
		return 400, "Debe especificar la descripcion de la direccion"
	}
	if len(t.DirecTelefono) == 0 {
		return 400, "Debe especificar el telefono de la direccion"
	}
	if len(t.DirecCorreo) == 0 {
		return 400, "Debe especificar el correo de la direccion"
	}

	result, err := bd.InsertDireccion(t)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el registro de la direccion " + t.DirecNombre + " > " + err.Error()
	}
	return 200, "{ DireccionID: " + result + "}"

}

func SelectDireccion(body string, request events.APIGatewayV2HTTPRequest) (int, string) {
	var err error
	var id_direc int
	var Slug string

	if len(request.QueryStringParameters["id_direccion"]) > 0 {
		id_direc, err = strconv.Atoi(request.QueryStringParameters["id_direccion"])
		if err != nil {
			return 500, "Ocurrio un error al intentar convertir en entero el valor " + request.QueryStringParameters["id_direccion"]
		}
	} else {
		if len(request.QueryStringParameters["slug"]) > 0 {
			Slug = request.QueryStringParameters["slug"]
		}
	}

	lista, err2 := bd.SelectDireccion(id_direc, Slug)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar capturar Direccion/s > " + err2.Error()
	}

	Direc, err3 := json.Marshal(lista)
	if err3 != nil {
		return 400, "Ocurrio un error al intentar convertir en JSON Direccion/s > " + err3.Error()
	}

	return 200, string(Direc)
}

func DeleteFisicoDireccion(body string, User string, id int) (int, string) {
	if id == 0 {
		return 400, "Debe especificar ID de la Direccion a Borrar"
	}

	err := bd.DeleteFisicoDireccion(id)
	if err != nil {
		return 400, "Ocurrió un error al intentar realizar el DELETE de la direccion " + strconv.Itoa(id) + " > " + err.Error()
	}

	return 200, "Delete OK"
}

func UpdateDireccion(body string, User string, id int) (int, string) {
	var t models.Direccion

	err := json.Unmarshal([]byte(body), &t)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	t.DirecID = id
	err2 := bd.UpdateDireccion(t)
	if err2 != nil {
		return 400, "Ocurrio un error al intentar realizar el UPDATE de la Direccion " + strconv.Itoa(id) + " > " + err2.Error()
	}

	return 200, "Update OK"
}
