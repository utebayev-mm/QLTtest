package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func (t *Category) DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./static/templates/deletecategory.html")
	if err != nil {
		log.Println(err)
	}
	if r.Method == http.MethodPost {
		CategoryID := r.FormValue("CategoryIDToDelete")
		ID, errID := strconv.Atoi(CategoryID)
		if errID != nil {
			log.Println(errID)
		}
		t.repo.DeleteCategory(ID)
		http.Redirect(w, r, "/browsecategories", http.StatusSeeOther)

	}
	errExec := template.Execute(w, "")
	if errExec != nil {
		log.Println(errExec)
	}
}

func (t *Category) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	IDtoDelete, err := strconv.Atoi(r.URL.Path[16:])
	fmt.Println(IDtoDelete)
	if err != nil {
		log.Println(err)
	}
	t.repo.DeleteCategory(IDtoDelete)
	http.Redirect(w, r, "/browsecategories", http.StatusSeeOther)
}
