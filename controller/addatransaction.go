package controller

import (
	"fmt"
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

func NewTransaction(repo *utils.Repo) *Transaction {
	return &Transaction{repo: repo}
}

func (t *Transaction) AddATransaction(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./static/templates/addatransaction.html")
	if err != nil {
		log.Println(err)
	}
	if r.Method == http.MethodPost {
		TransactionName := r.FormValue("Transactionname")
		TransactionValue := r.FormValue("Transactionvalue")
		TransactionType := r.FormValue("Transactiontype")
		TransactionCategory := r.FormValue("Transactioncategory")
		TransactionComment := r.FormValue("Transactioncomment")
		fmt.Println(TransactionName)
		fmt.Println(TransactionValue)
		fmt.Println(TransactionType)
		fmt.Println(TransactionCategory)
		fmt.Println(TransactionComment)
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
		} else if TransactionType == "Expenditure" {
			Transaction.Type = false
		}
		Transaction.CategoryID = "1"
		Transaction.Comments = TransactionComment
		t.repo.Create(Transaction)
		http.Redirect(w, r, "/browsetransactions", http.StatusSeeOther)

	}
	errExec := template.Execute(w, "")
	if errExec != nil {
		log.Println(errExec)
	}
}
