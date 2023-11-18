package controller

import (
	hangman "hangman/Hangman"
	initTemplate "hangman/templates"
	"net/http"
)

func DataJeux(w http.ResponseWriter, r *http.Request) {
	hangman.Ecriremot()
	hangman.User.MotaDeviner = hangman.LireFichierMot()

	initTemplate.Temp.ExecuteTemplate(w, "jeux", hangman.User)
}

func TreatmentJeux(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/jeux", http.StatusSeeOther)
	}

	http.Redirect(w, r, "/jeux", http.StatusSeeOther)

}
