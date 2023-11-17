package controller

import (
	initTemplate "hangman/templates"
	"net/http"
)

func DataIndex(w http.ResponseWriter, r *http.Request) {

	initTemplate.Temp.ExecuteTemplate(w, "index", nil)
}

func TreatmentIndex(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Redirect(w, r, "/index", http.StatusSeeOther)
	}

	username := Name{r.FormValue("nom")}
	Nom = username
	http.Redirect(w, r, "/accueil", http.StatusSeeOther)

}
