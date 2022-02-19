package controller

import (
	"log"
	"net/http"
	"qltTestApi/model"
	"text/template"
)

func (t *Category) BrowseCategories(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./static/templates/browsecategories.html")
	if err != nil {
		log.Println(err)
	}
	var Categories []model.Category

	Categories = t.repo.BrowseCategories()

	errExec := template.Execute(w, Categories)
	if errExec != nil {
		log.Println(errExec)
	}
}
