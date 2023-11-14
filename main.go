package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// / STRUCT DATA
type UserData struct {
	Nom string
}

var username UserData

func main() {
	///// PARSER TEMPLATES HTML

	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println(fmt.Sprint("ERREUR => %s", err.Error()))
		return
	}

	//// ROUTE ACCUEIL JEUX
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := 0
		temp.ExecuteTemplate(w, "index", data)
	})

	/// ROUTE TREATMENT PSEUDO JOUEUR
	http.HandleFunc("/treatmentlogin", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			username = UserData{
				Nom: r.FormValue("nom"),
			}
			http.Redirect(w, r, "/accueil", http.StatusSeeOther)
		}
		temp.ExecuteTemplate(w, "treatmentlogin", username)
	})

	// / TEST
	http.HandleFunc("/accueil", func(w http.ResponseWriter, r *http.Request) {
		datadisplay := username
		temp.ExecuteTemplate(w, "accueil", datadisplay)
	})

	//// INIT DU SERVEUR LOCAL
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)
}
