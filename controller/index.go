package controller

import (
	hangman "hangman/Hangman"
	initTemplate "hangman/templates"
	"net/http"
)

func DataIndex(w http.ResponseWriter, r *http.Request) {

	if hangman.User.UserURL != "" {
		http.Redirect(w, r, "/index", http.StatusPermanentRedirect)
	}
	initTemplate.Temp.ExecuteTemplate(w, "index", nil)
}

func TreatmentIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/index", http.StatusPermanentRedirect)
	}
	hangman.User.Username = r.FormValue("nom")
	hangman.User.UserURL = "/accueil"
	http.Redirect(w, r, "/accueil", http.StatusPermanentRedirect)

}
