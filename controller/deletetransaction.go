package controller

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

func (t *Transaction) DeleteTransactionByID(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./static/templates/deletetransaction.html")
	if err != nil {
		log.Println(err)
	}
	if r.Method == http.MethodPost {
		TransactionID := r.FormValue("TransactionIDToDelete")
		ID, errID := strconv.Atoi(TransactionID)
		if errID != nil {
			log.Println(errID)
		}
		t.repo.Delete(ID)
		http.Redirect(w, r, "/browsetransactions", http.StatusSeeOther)

	}
	errExec := template.Execute(w, "")
	if errExec != nil {
		log.Println(errExec)
	}
}

func (t *Transaction) DeleteTransaction(w http.ResponseWriter, r *http.Request) {
	IDtoDelete, err := strconv.Atoi(r.URL.Path[19:])
	fmt.Println(IDtoDelete)
	if err != nil {
		log.Println(err)
	}
	t.repo.Delete(IDtoDelete)
	http.Redirect(w, r, "/browsetransactions", http.StatusSeeOther)
}
