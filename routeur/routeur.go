package routeur

import (
	"hangman/controller"
	"net/http"
	"os"
)

func InitServer() {
	http.HandleFunc("/index", controller.DataIndex)
	http.HandleFunc("/accueil", controller.DataAccueil)
	http.HandleFunc("/treatmentlogin", controller.TreatmentIndex)
	http.HandleFunc("/treatmentaccueil", controller.TreatmentAccueil)
	http.HandleFunc("/jeux", controller.DataJeux)
	http.HandleFunc("/treatmentjeux", controller.TreatmentJeux)
	http.HandleFunc("/resultat", controller.DataResultat)
	http.HandleFunc("/treatmentresultat", controller.TreatmentResultat)

	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)
}
