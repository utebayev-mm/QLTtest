package controller

import (
	"fmt"
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
	template, err := template.ParseFiles("./static/templates/addacategory.html")
	if err != nil {
		log.Println(err)
	}
	if r.Method == http.MethodPost {
		CategoryName := r.FormValue("Categoryname")
		fmt.Println(CategoryName)
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
