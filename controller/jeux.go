package controller

import (
	hangman "hangman/Hangman"
	initTemplate "hangman/templates"
	"net/http"
)

func DataJeux(w http.ResponseWriter, r *http.Request) {

	hangman.Ecriremot()
	Word := hangman.LireFichierMot()

	datajeux := User{Word}

	initTemplate.Temp.ExecuteTemplate(w, "jeux", datajeux)

}
