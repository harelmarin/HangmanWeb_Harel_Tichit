package controller

import (
	hangman "hangman/Hangman"
	initTemplate "hangman/templates"
	"net/http"
)

func DataResultat(w http.ResponseWriter, r *http.Request) {
	if hangman.User.UserURL != "/resultat" {
		http.Redirect(w, r, hangman.User.UserURL, http.StatusPermanentRedirect)
	}
	initTemplate.Temp.ExecuteTemplate(w, "resultat", hangman.User)
}

func TreatmentResultat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/resultat", http.StatusPermanentRedirect)
	}
	hangman.User.UserURL = ""
	http.Redirect(w, r, "/index", http.StatusPermanentRedirect)
}
