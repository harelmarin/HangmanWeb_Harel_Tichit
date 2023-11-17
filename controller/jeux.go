package controller

import (
	initTemplate "hangman/templates"
	"net/http"
)

func DataJeux(w http.ResponseWriter, r *http.Request) {
	datajeux := User{DataUser.Username, DataUser.Difficulte}
	initTemplate.Temp.ExecuteTemplate(w, "jeux", datajeux)
}
