package controller

import (
	initTemplate "hangman/templates"
	"net/http"
)

func DataAccueil(w http.ResponseWriter, r *http.Request) {
	dataaccueil := User{DataUser.Username, DataUser.Difficulte}
	initTemplate.Temp.ExecuteTemplate(w, "accueil", dataaccueil)
}

func TreatmentAccueil(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/accueil", http.StatusSeeOther)
	}
	user := User{r.FormValue("nom"), r.FormValue("difficulte")}
	DataUser = user
	http.Redirect(w, r, "/jeux", http.StatusSeeOther)
}
