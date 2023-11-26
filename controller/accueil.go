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
	hangman.User.Ecriremot()
	hangman.User.MotaDeviner = hangman.LireFichierMot()
	if hangman.User.Difficulte == "Facile" {
		lettrealeatoire1 := hangman.User.ChoisirLettreAleatoire(hangman.User.MotaDeviner)
		lettrealeatoire2 := hangman.User.ChoisirLettreAleatoire(hangman.User.MotaDeviner)
		if lettrealeatoire2 == lettrealeatoire1 {
			lettrealeatoire2 = hangman.User.ChoisirLettreAleatoire(hangman.User.MotaDeviner)
		}
		hangman.User.LettreDejaJoue = append(hangman.User.LettreDejaJoue, lettrealeatoire1)
		hangman.User.LettreDejaTrouve = append(hangman.User.LettreDejaTrouve, lettrealeatoire1)
		hangman.User.LettreDejaJoue = append(hangman.User.LettreDejaJoue, lettrealeatoire2)
		hangman.User.LettreDejaTrouve = append(hangman.User.LettreDejaTrouve, lettrealeatoire2)
	} else if hangman.User.Difficulte == "Moyen" {
		lettrealeatoire := hangman.User.ChoisirLettreAleatoire(hangman.User.MotaDeviner)
		hangman.User.LettreDejaJoue = append(hangman.User.LettreDejaJoue, lettrealeatoire)
		hangman.User.LettreDejaTrouve = append(hangman.User.LettreDejaTrouve, lettrealeatoire)
	} else if hangman.User.Difficulte == "Difficile" {

	}
	hangman.User.TentativesRestantes = 9
	hangman.User.UserURL = "/jeux"
	http.Redirect(w, r, "/jeux", http.StatusPermanentRedirect)
}
