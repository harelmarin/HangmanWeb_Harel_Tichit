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
	hangman.User.LettreDejaJoue = []string{}
	hangman.User.LettreDejaTrouve = []string{}
	hangman.User.MotaDeviner = ""
	hangman.User.Difficulte = r.FormValue("difficulte")
	hangman.Ecriremot()
	hangman.User.TentativesRestantes = 9
	hangman.User.MotaDeviner = hangman.LireFichierMot()
	http.Redirect(w, r, "/jeux", http.StatusSeeOther)
}
