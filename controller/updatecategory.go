package controller

import (
	"log"
	"net/http"
	"qltTestApi/model"
	"strconv"
	"text/template"
)

func (t *Category) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./static/templates/updatecategory.html")
	if err != nil {
		log.Println(err)
	}
	IDtoEdit, err := strconv.Atoi(r.URL.Path[16:])
	if err != nil {
		log.Println(err)
	}
	CategoryToEdit := t.repo.SelectCategory(IDtoEdit)
	if r.Method == http.MethodPost {
		CategoryName := r.FormValue("Categoryname")
		var Category model.Category
		Category.Name = CategoryName
		Category.ID = IDtoEdit
		t.repo.UpdateCategory(Category)
		http.Redirect(w, r, "/browsecategories", http.StatusSeeOther)
	}
	errExec := template.Execute(w, CategoryToEdit)
	if errExec != nil {
		log.Println(errExec)
	}
}
