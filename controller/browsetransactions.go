package controller

import (
	"log"
	"net/http"
	"qltTestApi/model"
	"text/template"
)

type TransactionBrowseData struct {
	Categories   []model.Category
	Transactions []model.Transaction
}

func (t *Transaction) BrowseTransactions(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./static/templates/browsetransactions.html")
	if err != nil {
		log.Println(err)
	}
	var Transactions []model.Transaction

	Transactions = t.repo.Browse()
	var BrowseData TransactionBrowseData
	var Categories []model.Category
	Categories = t.repo.BrowseCategories()
	BrowseData.Transactions = Transactions
	BrowseData.Categories = Categories
	errExec := template.Execute(w, BrowseData)
	if errExec != nil {
		log.Println(errExec)
	}
}
