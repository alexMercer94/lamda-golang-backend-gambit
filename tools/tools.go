package tools

import (
	"fmt"
	"strings"
	"time"
)

/*
Formatear fecha para guardar en MySQL
*/
func DateMySQL() string {
	t := time.Now()
	return fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute(), t.Second())
}

/*
* Eliminar caracteres de un string
 */
func EscapeString(str string) string {
	desc := strings.ReplaceAll(str, "'", "")
	desc = strings.ReplaceAll(desc, "\"", "")

	return desc
}
