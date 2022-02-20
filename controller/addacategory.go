package controller

import (
	"log"
	"net/http"
	"qltTestApi/model"
	"qltTestApi/utils"
	"text/template"
)

type Category struct {
	repo *utils.Repo
}

func NewCategory(repo *utils.Repo) *Category {
	return &Category{repo: repo}
}

func (t *Category) AddACategory(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"./static/templates/addacategory.html",
		"./static/templates/footer.html",
	}
	template, err := template.ParseFiles(templates...)
	if err != nil {
		log.Println(err)
	}
	if r.Method == http.MethodPost {
		CategoryName := r.FormValue("Categoryname")
		var Category model.Category
		Category.Name = CategoryName
		t.repo.CreateCategory(Category)
		http.Redirect(w, r, "/browsecategories", http.StatusSeeOther)
	}
	errExec := template.Execute(w, "")
	if errExec != nil {
		log.Println(errExec)
	}
}
