package controller

import (
	hangman "hangman/Hangman"
	initTemplate "hangman/templates"
	"net/http"
	"regexp"
	"strings"
)

func DataJeux(w http.ResponseWriter, r *http.Request) {
	if len(hangman.User.LettreJoueur) == 1 {
		if hangman.User.LettreDejaChoisie(hangman.User.LettreJoueur) {
			hangman.User.MessageEnvoi = "Vous avez déjà émis cette lettre !"
		} else if strings.Contains(hangman.User.MotaDeviner, hangman.User.LettreJoueur) {

			hangman.User.LettreDejaJoue = append(hangman.User.LettreDejaJoue, hangman.User.LettreJoueur)
			hangman.User.MessageEnvoi = "Vous avez trouvé une lettre !"
			hangman.User.LettreDejaTrouve = append(hangman.User.LettreDejaTrouve, hangman.User.LettreJoueur)
		} else {
			hangman.User.LettreDejaJoue = append(hangman.User.LettreDejaJoue, hangman.User.LettreJoueur)
			hangman.User.MessageEnvoi = "La lettre n'est pas dans le mot !"
			hangman.User.TentativesRestantes--
		}
	}
	hangman.User.MotCache = hangman.User.AffichageTirets()
	initTemplate.Temp.ExecuteTemplate(w, "jeux", hangman.User)
}

func TreatmentJeux(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/index", http.StatusMovedPermanently)
	}
	hangman.User.LettreJoueur = r.FormValue("lettre")
	CheckValeur, _ := regexp.MatchString("^[a-zA-Z]{1,25}$", hangman.User.LettreJoueur)
	if !CheckValeur {
		hangman.User.MessageEnvoi = "Invalide"
		hangman.User.LettreJoueur = ""
	}
	hangman.User.CheckLettreJoueur = false
	http.Redirect(w, r, "/jeux", http.StatusMovedPermanently)
}
