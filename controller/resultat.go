package controller

import (
	hangman "hangman/Hangman"
	initTemplate "hangman/templates"
	"net/http"
)

func DataResultat(w http.ResponseWriter, r *http.Request) {

	initTemplate.Temp.ExecuteTemplate(w, "resultat", hangman.User)
}

func TreatmentResultat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/resultat", http.StatusSeeOther)
	}

	http.Redirect(w, r, "/accueil", http.StatusSeeOther)

}
