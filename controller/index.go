package controller

import (
	"log"
	"net/http"
	"text/template"
)

func MainPage(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"./static/templates/index.html",
		"./static/templates/footer.html",
	}
	template, err := template.ParseFiles(templates...)
	if err != nil {
		log.Println(err)
	}
	errExec := template.Execute(w, "")
	if errExec != nil {
		log.Println(errExec)
	}
}
