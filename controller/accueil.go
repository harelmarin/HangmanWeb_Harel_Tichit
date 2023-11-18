package controller

import (
	hangman "hangman/Hangman"
	initTemplate "hangman/templates"
	"net/http"
)

func DataAccueil(w http.ResponseWriter, r *http.Request) {

	initTemplate.Temp.ExecuteTemplate(w, "accueil", hangman.User)
}

func TreatmentAccueil(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/accueil", http.StatusSeeOther)
	}
	hangman.User.Difficulte = r.FormValue("difficulte")
	http.Redirect(w, r, "/jeux", http.StatusSeeOther)
}
