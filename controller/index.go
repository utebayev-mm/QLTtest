package controller

import (
	"log"
	"net/http"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./static/templates/index.html")
	if err != nil {
		log.Println(err)
	}
	errExec := template.Execute(w, "")
	if errExec != nil {
		log.Println(errExec)
	}
}
