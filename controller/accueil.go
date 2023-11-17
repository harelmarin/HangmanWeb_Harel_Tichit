package controller

import (
	initTemplate "hangman/templates"
	"net/http"
)

func DataAccueil(w http.ResponseWriter, r *http.Request) {
	dataaccueil := Name{Nom.Username}
	initTemplate.Temp.ExecuteTemplate(w, "accueil", dataaccueil)
}

func TreatmentAccueil(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/accueil", http.StatusSeeOther)
	}
	userdiff := Difficulty{r.FormValue("difficulte")}
	Diff = userdiff
	http.Redirect(w, r, "/jeux", http.StatusSeeOther)
}
