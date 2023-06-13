package db

import (
	awsecretm "backend/lamda-golang-backend-gambit/awssecretm"
	"backend/lamda-golang-backend-gambit/models"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var SecretModel models.SecretRDSJson
var err error
var Db *sql.DB

/*
Leer Secret de SecretManager
*/
func ReadSecret() error {
	SecretModel, err = awsecretm.GetSecret(os.Getenv("SecretName"))
	return err
}

/*
Conectar a la Base de Datos
*/
func DbConnect() error {
	Db, err = sql.Open("mysql", ConnStr(SecretModel))

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = Db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	fmt.Println("Conexion exitosa a la BD")
	return nil
}

/*
Generar string de conexion a la BD
*/
func ConnStr(key models.SecretRDSJson) string {
	var dbUser, authToken, dbEndpoint, dbName string
	dbUser = key.Username
	authToken = key.Password
	dbEndpoint = key.Host
	dbName = "gambit"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?allowCleartextPasswords=true", dbUser, authToken, dbEndpoint, dbName)
	fmt.Println(dsn)
	return dsn
}
