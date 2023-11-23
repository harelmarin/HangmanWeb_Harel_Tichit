package controller

import (
	hangman "hangman/Hangman"
	initTemplate "hangman/templates"
	"net/http"
	"regexp"
	"strings"
)

func DataJeux(w http.ResponseWriter, r *http.Request) {
	if hangman.User.UserURL != "/jeux" {
		http.Redirect(w, r, hangman.User.UserURL, http.StatusPermanentRedirect)
	}
	hangman.User.LettreJoueur = strings.ToLower(hangman.User.LettreJoueur)
	if len(hangman.User.LettreJoueur) == 1 {
		if hangman.User.LettreDejaChoisie(hangman.User.LettreJoueur) {
			hangman.User.MessageEnvoi = "Vous avez déjà émis cette lettre !"
		} else if strings.Contains(hangman.User.MotaDeviner, hangman.User.LettreJoueur) {
			hangman.User.LettreDejaJoue = append(hangman.User.LettreDejaJoue, hangman.User.LettreJoueur)
			hangman.User.MessageEnvoi = "Vous avez trouvé une lettre !"
			hangman.User.LettreDejaTrouve = append(hangman.User.LettreDejaTrouve, hangman.User.LettreJoueur)
			if hangman.User.ToutesLettresTrouvees(hangman.User.MotaDeviner, hangman.User.LettreDejaTrouve) {
				hangman.User.MessageEnvoi = "Vous avez trouvé toutes les lettres !"
				hangman.User.CheckVictoire = true
				hangman.User.UserURL = "/resultat"
				http.Redirect(w, r, "/resultat", http.StatusPermanentRedirect)
			}
		} else {
			hangman.User.LettreDejaJoue = append(hangman.User.LettreDejaJoue, hangman.User.LettreJoueur)
			hangman.User.MessageEnvoi = "La lettre n'est pas dans le mot !"
			hangman.User.TentativesRestantes--
			if hangman.User.TentativesRestantes <= 0 {
				hangman.User.CheckVictoire = false
				hangman.User.UserURL = "/resultat"
				http.Redirect(w, r, "/resultat", http.StatusPermanentRedirect)
			}
		}
	}
	if len(hangman.User.LettreJoueur) > 1 {
		if hangman.User.MotaDeviner == hangman.User.LettreJoueur {
			hangman.User.CheckVictoire = true
			hangman.User.UserURL = "/resultat"
			http.Redirect(w, r, "/resultat", http.StatusPermanentRedirect)
		} else {
			hangman.User.MessageEnvoi = "Ce n'est pas le bon mot !"
			hangman.User.TentativesRestantes -= 2
			if hangman.User.TentativesRestantes <= 0 {
				hangman.User.CheckVictoire = false
				hangman.User.UserURL = "/resultat"
				http.Redirect(w, r, "/resultat", http.StatusPermanentRedirect)
			}
		}
	}
	hangman.User.MotCache = hangman.User.AffichageTirets()
	initTemplate.Temp.ExecuteTemplate(w, "jeux", hangman.User)
}

func TreatmentJeux(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/accueil", http.StatusPermanentRedirect)
	}
	hangman.User.LettreJoueur = r.FormValue("lettre")
	CheckValeur, _ := regexp.MatchString("^[a-zA-Z]{1,15}$", hangman.User.LettreJoueur)
	if !CheckValeur {
		hangman.User.MessageEnvoi = "Invalide"
		hangman.User.LettreJoueur = ""
	}
	hangman.User.CheckLettreJoueur = false
	http.Redirect(w, r, "/jeux", http.StatusPermanentRedirect)
}
