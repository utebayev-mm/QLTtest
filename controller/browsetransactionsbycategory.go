package controller

import (
	"log"
	"net/http"
	"qltTestApi/model"
	"strconv"
	"text/template"
)

func (t *Transaction) BrowseTransactionsByCategory(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"./static/templates/browsebycategory.html",
		"./static/templates/footer.html",
	}
	template, err := template.ParseFiles(templates...)
	if err != nil {
		log.Println(err)
	}
	var Transactions []model.Transaction
	CategoryID, err := strconv.Atoi(r.URL.Path[18:])
	if err != nil {
		log.Println(err)
	}
	Transactions = t.repo.BrowseByCategory(CategoryID)
	var BrowseData TransactionBrowseData
	var Categories []model.Category
	Categories = t.repo.BrowseCategories()
	BrowseData.Transactions = Transactions
	BrowseData.Categories = Categories
	if len(Transactions) > 0 {
		BrowseData.Info = Transactions[0].CategoryName
	}
	errExec := template.Execute(w, BrowseData)
	if errExec != nil {
		log.Println(errExec)
	}
}
