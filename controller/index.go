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

	user := User{r.FormValue("nom"), r.FormValue("difficulte")}
	DataUser = user
	http.Redirect(w, r, "/accueil", http.StatusSeeOther)

}
