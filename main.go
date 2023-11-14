package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

func main() {
	///// PARSER TEMPLATES HTML

	temp, err := template.ParseGlob("./templates/*.html")
	if err != nil {
		fmt.Println(fmt.Sprint("ERREUR => %s", err.Error()))
		return
	}

	//// 1ERE ROUTE
	http.HandleFunc("/index", func(w http.ResponseWriter, r *http.Request) {
		data := 0
		temp.ExecuteTemplate(w, "index", data)
	})

	//// INIT DU SERVEUR LOCAL
	rootDoc, _ := os.Getwd()
	fileserver := http.FileServer(http.Dir(rootDoc + "/asset"))
	http.Handle("/static/", http.StripPrefix("/static/", fileserver))

	http.ListenAndServe("localhost:8080", nil)
}
