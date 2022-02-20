package controller

import (
	"log"
	"net/http"
	"qltTestApi/model"
	"text/template"
)

func (t *Category) BrowseCategories(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"./static/templates/browsecategories.html",
		"./static/templates/footer.html",
	}
	template, err := template.ParseFiles(templates...)
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
