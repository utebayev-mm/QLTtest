package controller

import (
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func (t *Category) DeleteCategoryByID(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"./static/templates/deletecategory.html",
		"./static/templates/footer.html",
	}
	template, err := template.ParseFiles(templates...)
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
	if err != nil {
		log.Println(err)
	}
	t.repo.DeleteCategory(IDtoDelete)
	http.Redirect(w, r, "/browsecategories", http.StatusSeeOther)
}
