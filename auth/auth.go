package auth

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type TokenJSON struct {
	Sub       string
	Event_Id  string
	Token_Use string
	Scope     string
	Auth_time int
	Iss       string
	Exp       int
	Iat       int
	Client_id string
	Username  string
}

/*
Validar un token de JWT
*/
func ValidateToken(token string) (bool, error, string) {
	parts := strings.Split(token, ".")

	if len(parts) != 3 {
		fmt.Println("El token no es válido")
		return false, nil, "El token no es válido"
	}

	// Decodificar el payload del token JWT
	userInfo, err := base64.StdEncoding.DecodeString(parts[1])

	if err != nil {
		fmt.Println("No se puede decoficar la parte del token: ", err.Error())
		return false, err, err.Error()
	}

	// Convertir en el JSON en una estructura con los datos
	var tkj TokenJSON
	err = json.Unmarshal(userInfo, &tkj)

	if err != nil {
		fmt.Println("No se puede decodificar la estructura JSON ", err.Error())
		return false, err, err.Error()
	}

	// Verificar que el token no este expirado
	now := time.Now()
	tm := time.Unix(int64(tkj.Exp), 0)

	if tm.Before(now) {
		fmt.Println("Fecha de expiración del token: " + tm.String())
		fmt.Println(" > Token expirado")

		return false, err, "Token expirado"
	}

	return true, nil, string(tkj.Username)
}
