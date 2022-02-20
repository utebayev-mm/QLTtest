package controller

import (
	"log"
	"net/http"
	"qltTestApi/model"
	"qltTestApi/utils"
	"strconv"
	"text/template"
	"time"
)

type Transaction struct {
	repo *utils.Repo
}

type TransactionData struct {
	Categories []model.Category
}

func NewTransaction(repo *utils.Repo) *Transaction {
	return &Transaction{repo: repo}
}

func (t *Transaction) AddATransaction(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"./static/templates/addatransaction.html",
		"./static/templates/footer.html",
	}
	template, err := template.ParseFiles(templates...)
	if err != nil {
		log.Println(err)
	}
	var TransactionDataCategories TransactionData
	if r.Method == http.MethodPost {
		TransactionName := r.FormValue("Transactionname")
		TransactionValue := r.FormValue("Transactionvalue")
		TransactionType := r.FormValue("Transactiontype")
		TransactionCategory := r.FormValue("Transactioncategory")
		TransactionComment := r.FormValue("Transactioncomment")
		var Transaction model.Transaction
		Transaction.Name = TransactionName
		Transaction.Price, err = strconv.Atoi(TransactionValue)
		if err != nil {
			log.Println(err)
		}
		dt := time.Now()
		Transaction.Date = dt.Format("01-02-2006 15:04:05")
		if TransactionType == "Income" {
			Transaction.Type = true
			if Transaction.Price < 0 {
				Transaction.Price = Transaction.Price * -1
			}
		} else if TransactionType == "Expenditure" {
			Transaction.Type = false
			if Transaction.Price > 0 {
				Transaction.Price = Transaction.Price * -1
			}
		}
		Category := t.repo.SelectCategorybyName(TransactionCategory)
		Transaction.CategoryName = Category.Name
		Transaction.CategoryID = Category.ID
		Transaction.Comments = TransactionComment
		Transaction.Categories = t.repo.BrowseCategories()

		t.repo.Create(Transaction)
		http.Redirect(w, r, "/browsetransactions", http.StatusSeeOther)

	}
	TransactionDataCategories.Categories = t.repo.BrowseCategories()
	errExec := template.Execute(w, TransactionDataCategories)
	if errExec != nil {
		log.Println(errExec)
	}
}
