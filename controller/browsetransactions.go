package controller

import (
	"log"
	"net/http"
	"qltTestApi/model"
	"text/template"
)

func (t *Transaction) BrowseTransactions(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./static/templates/browsetransactions.html")
	if err != nil {
		log.Println(err)
	}
	var Transactions []model.Transaction

	Transactions = t.repo.Browse()

	errExec := template.Execute(w, Transactions)
	if errExec != nil {
		log.Println(errExec)
	}
}
