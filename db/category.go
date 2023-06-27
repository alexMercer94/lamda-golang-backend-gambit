package db

import (
	"database/sql"
	"fmt"

	/* "strconv"
	"strings" */

	"backend/lamda-golang-backend-gambit/models"

	_ "github.com/go-sql-driver/mysql"
	// "backend/lamda-golang-backend-gambit/tools"
)

/*
* Operación de Inserción de la BD en la tabla category
 */
func InsertCategory(cat models.Category) (int64, error) {
	fmt.Println("Comienza Operación de InsertCategory")

	err := DbConnect()
	if err != nil {
		return 0, err
	}
	defer Db.Close()

	sqlQuery := fmt.Sprintf(`INSERT INTO category (Categ_Name, Categ_Path) VALUES ('%v', '%v')`, cat.CategName, cat.CategPath)

	var result sql.Result
	result, err = Db.Exec(sqlQuery)

	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}

	lastInsertId, resulErr := result.LastInsertId()
	if resulErr != nil {
		return 0, resulErr
	}

	fmt.Println("Operacion de InsertCategory > Success")
	return lastInsertId, nil
}
