package initTemplate

import (
	"fmt"
	"html/template"
	"os"
)

var Temp *template.Template

// /// PARSER TEMPLATES HTML
func InitTemplate() {
	temp, errTemp := template.ParseGlob("./templates/*.html")
	if errTemp != nil {
		fmt.Println(fmt.Sprint("ERREUR => %s", errTemp.Error()))
		os.Exit(1)
	}
	Temp = temp
}
