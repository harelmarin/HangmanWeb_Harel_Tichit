package controller

import (
	initTemplate "hangman/templates"
	"net/http"
)

func DataJeux(w http.ResponseWriter, r *http.Request) {
	datajeux := Difficulty{Diff.Difficulte}

	initTemplate.Temp.ExecuteTemplate(w, "jeux", datajeux)

}
