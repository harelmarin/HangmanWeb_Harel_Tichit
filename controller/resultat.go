package controller

import (
	initTemplate "hangman/templates"
	"net/http"
)

func DataResultat(w http.ResponseWriter, r *http.Request) {

	initTemplate.Temp.ExecuteTemplate(w, "resultat", nil)
}
