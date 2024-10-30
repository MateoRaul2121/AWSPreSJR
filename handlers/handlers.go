package handlers

import (
	"fmt"
	"strconv"

	"example.com/m/v2/auth"
	"example.com/m/v2/routers"
	"github.com/aws/aws-lambda-go/events"
)

func Manejadores(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {

	fmt.Println("Voy a procesar " + path + " > " + method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := validoAuthorization(path, method, headers)
	if !isOk {
		return statusCode, user
	}

	fmt.Println("path[0:4] = " + path[0:4])

	switch path[0:4] {
	case "user":
		return ProcesoUsers(body, path, method, user, id, request)
	case "jefa":
		return ProcesoJefatura(body, path, method, user, idn, request)
	case "dire":
		return ProcesoDireccion(body, path, method, user, idn, request)
	case "secr":
		return ProcesoSecretaria(body, path, method, user, idn, request)
	}
	return 400, "Method invalid"
}

func validoAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") ||
		(path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido"
	}

	todoOK, err, msg := auth.ValidoToken(token)
	if !todoOK {
		if err != nil {
			fmt.Println("Error en el token " + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error en el token " + msg)
			return false, 401, msg
		}
	}

	fmt.Println("Token OK")
	return true, 200, msg
}

func ProcesoUsers(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func ProcesoJefatura(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertJefatura(body, user)
	case "GET":
		return routers.SelectJefatura(body, request)
	case "DELETE":
		return routers.DeleteFisicoJefatura(body, user, id)
	case "PUT":
		return routers.UpdateJefatura(body, user, id)
	}
	return 400, "Method Invalid"
}

func ProcesoDireccion(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertDireccion(body, user)
	case "GET":
		return routers.SelectDireccion(body, request)
	case "DELETE":
		return routers.DeleteFisicoDireccion(body, user, id)
	case "PUT":
		return routers.UpdateDireccion(body, user, id)
	}
	return 400, "Method Invalid"
}

func ProcesoSecretaria(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertSecretaria(body, user)
	case "GET":
		return routers.SelectSecretaria(body, request)
	case "DELETE":
		return routers.DeleteFisicoSecretaria(body, user, id)
	case "PUT":
		return routers.UpdateSecretaria(body, user, id)
	}

	return 400, "Method Invalid"
}
