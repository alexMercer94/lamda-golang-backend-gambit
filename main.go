package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"

	"backend/lamda-golang-backend-gambit/awsgo"
	"backend/lamda-golang-backend-gambit/db"
	"backend/lamda-golang-backend-gambit/handlers"
	"os"

	lambda "github.com/aws/aws-lambda-go/lambda"
)

/*
Ejecucion de la lambda
*/
func main() {
	lambda.Start(ExecuteLambda)
}

/*
Logica de la lambda
*/
func ExecuteLambda(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	// Se inicializa AWS
	awsgo.InitializeAWS()

	// Se validan las variables de entorno
	if !ValidateParams() {
		panic("Error en los parametros; debe enviar 'SecretName', 'UrlPrefix'")
	}

	// Se obtienen los valores para los Handlers
	var res *events.APIGatewayProxyResponse
	path := strings.Replace(request.RequestContext.ResourcePath, os.Getenv("UrlPrefix"), "", -1)
	method := request.RequestContext.HTTPMethod
	body := request.Body
	headers := request.Headers

	// Se obtiene el valor del secreto de SecretManager
	db.ReadSecret()

	status, message := handlers.Handlers(path, method, body, headers, request)

	headersResp := map[string]string{
		"content-Type": "application/json",
	}

	res = &events.APIGatewayProxyResponse{
		StatusCode: status,
		Body:       string(message),
		Headers:    headersResp,
	}

	return res, nil
}

/*
Validar si las variables de entorno de la Lambda son devueltas por AWS
*/
func ValidateParams() bool {
	_, getParam := os.LookupEnv("SecretName")

	if !getParam {
		return getParam
	}

	_, getParam = os.LookupEnv("UrlPrefix")

	if !getParam {
		return getParam
	}

	return getParam
}
