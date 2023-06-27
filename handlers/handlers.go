package handlers

import (
	"fmt"
	"strconv"

	"backend/lamda-golang-backend-gambit/auth"
	"backend/lamda-golang-backend-gambit/routers"

	"github.com/aws/aws-lambda-go/events"
)

/*
Handlers para las Rutas de la API
*/
func Handlers(path string, method string, body string, headers map[string]string, request events.APIGatewayV2HTTPRequest) (int, string) {
	fmt.Println("Se va a procesar " + path + " > " + method)

	id := request.PathParameters["id"]
	idn, _ := strconv.Atoi(id)

	isOk, statusCode, user := validateAuthorization(path, method, headers)

	if !isOk {
		return statusCode, user
	}

	switch path[0:4] {
	case "user":
		return UsersProcess(body, path, method, user, id, request)
	case "prod":
		return ProductsProcess(body, path, method, user, idn, request)
	case "stoc":
		return StockProcess(body, path, method, user, idn, request)
	case "addr":
		return AddressesProcess(body, path, method, user, idn, request)
	case "cate":
		return CategoryProcess(body, path, method, user, idn, request)
	case "orde":
		return OrdersProcess(body, path, method, user, idn, request)
	}

	return 400, "Method Invalid"
}

func UsersProcess(body string, path string, method string, user string, id string, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func ProductsProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func CategoryProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	switch method {
	case "POST":
		return routers.InsertCategory(body, user)
	}

	return 400, "Method Invalid"
}

func StockProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func AddressesProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

func OrdersProcess(body string, path string, method string, user string, id int, request events.APIGatewayV2HTTPRequest) (int, string) {
	return 400, "Method Invalid"
}

/*
Validar el token del Header Authorization de los request
*/
func validateAuthorization(path string, method string, headers map[string]string) (bool, int, string) {
	if (path == "product" && method == "GET") || (path == "category" && method == "GET") {
		return true, 200, ""
	}

	token := headers["authorization"]
	if len(token) == 0 {
		return false, 401, "Token requerido"
	}

	allOK, err, msg := auth.ValidateToken(token)
	if !allOK {
		if err != nil {
			fmt.Println("Error en el token" + err.Error())
			return false, 401, err.Error()
		} else {
			fmt.Println("Error en el token " + msg)
		}
	}

	fmt.Println("Token correcto")
	return true, 200, msg
}
