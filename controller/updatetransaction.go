package controller

import (
	"log"
	"net/http"
	"qltTestApi/model"
	"strconv"
	"text/template"
	"time"
)

func (t *Transaction) UpdateTransaction(w http.ResponseWriter, r *http.Request) {
	template, err := template.ParseFiles("./static/templates/updatetransaction.html")
	if err != nil {
		log.Println(err)
	}
	IDtoEdit, err := strconv.Atoi(r.URL.Path[19:])
	if err != nil {
		log.Println(err)
	}
	TransactionToEdit := t.repo.Select(IDtoEdit)
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
		} else if TransactionType == "Expenditure" {
			Transaction.Type = false
		}
		Transaction.CategoryID = TransactionCategory
		Transaction.Comments = TransactionComment
		Transaction.ID = IDtoEdit
		t.repo.Update(Transaction)
		http.Redirect(w, r, "/browsetransactions", http.StatusSeeOther)
	}
	errExec := template.Execute(w, TransactionToEdit)
	if errExec != nil {
		log.Println(errExec)
	}
}
