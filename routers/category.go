package routers

import (
	"encoding/json"
	"strconv"

	// "github.com/aws/aws-lambda-go/events"
	"backend/lamda-golang-backend-gambit/db"
	"backend/lamda-golang-backend-gambit/models"
)

func InsertCategory(body string, User string) (int, string) {
	var categoryModel models.Category
	// Setear informacion del body en JSON en el modelo de categoria
	err := json.Unmarshal([]byte(body), &categoryModel)
	if err != nil {
		return 400, "Error en los datos recibidos " + err.Error()
	}

	if len(categoryModel.CategName) == 0 {
		return 400, "Debe especificar el Nombre (Title) de la Categoría"
	}

	if len(categoryModel.CategPath) == 0 {
		return 400, "Debe especificar el Path (Ruta) de la Categoría"
	}

	isAdmin, msg := db.UserIsAdmin(User)
	if !isAdmin {
		return 400, msg
	}

	result, insertError := db.InsertCategory(categoryModel)
	if insertError != nil {
		return 400, "Ocurrio un error en el Insert del registro de la categoria " + categoryModel.CategName + " > " + insertError.Error()
	}

	return 200, "{ categId: " + strconv.Itoa(int(result)) + "}"
}
