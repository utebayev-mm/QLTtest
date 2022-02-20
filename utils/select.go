package utils

import (
	"log"
	"qltTestApi/model"
)

func (r Repo) Select(IDToEdit int) model.Transaction {
	var id, price int
	var name, date, comments, category_id string
	var transactionType bool
	err := r.db.QueryRow("SELECT id, name, price, date, type, comments ,category_id FROM transactions WHERE id = ?", IDToEdit).Scan(&id, &name, &price, &date, &transactionType, &comments, &category_id)
	if err != nil {
		log.Println(err)
	}
	Transaction := model.Transaction{
		ID:         id,
		Name:       name,
		Price:      price,
		Date:       date,
		Type:       transactionType,
		Comments:   comments,
		CategoryID: category_id,
	}
	return Transaction
}
