package controller

import (
	hangman "hangman/Hangman"
	initTemplate "hangman/templates"
	"net/http"
)

func DataAccueil(w http.ResponseWriter, r *http.Request) {
	if hangman.User.UserURL != "/accueil" {
		http.Redirect(w, r, hangman.User.UserURL, http.StatusPermanentRedirect)
	}
	initTemplate.Temp.ExecuteTemplate(w, "accueil", hangman.User)
}

func TreatmentAccueil(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/accueil", http.StatusPermanentRedirect)
	}
	hangman.User.LettreDejaJoue = []string{}
	hangman.User.LettreDejaTrouve = []string{}
	hangman.User.MessageEnvoi = "Choisissez une lettre ou un mot : "
	hangman.User.LettreJoueur = ""
	hangman.User.MotaDeviner = ""
	hangman.User.Difficulte = r.FormValue("difficulte") 
	hangman.Ecriremot()
	hangman.User.MotaDeviner = hangman.LireFichierMot()
	LettreAleatoire := hangman.User.ChoisirLettreAleatoire(hangman.User.MotaDeviner)
	hangman.User.LettreDejaJoue = append(hangman.User.LettreDejaJoue, LettreAleatoire)
	hangman.User.LettreDejaTrouve = append(hangman.User.LettreDejaTrouve, LettreAleatoire)
	hangman.User.ChoisirLettreAleatoire(hangman.User.MotaDeviner)
	hangman.User.TentativesRestantes = 9   
	hangman.User.UserURL = "/jeux"
	http.Redirect(w, r, "/jeux", http.StatusPermanentRedirect)
}
