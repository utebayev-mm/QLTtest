package utils

import (
	"fmt"
	"log"
	"qltTestApi/model"
)

func (r Repo) Browse() []model.Transaction {
	rows, err := r.db.Query("SELECT id, name, price, date, type, comments ,category_id FROM transactions")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var Transactions []model.Transaction
	for rows.Next() {
		var id, price int
		var name, date, comments, category_id string
		var transactionType bool
		rows.Scan(&id, &name, &price, &date, &transactionType, &comments, &category_id)
		transaction := model.Transaction{
			ID:         id,
			Name:       name,
			Price:      price,
			Date:       date,
			Type:       transactionType,
			Comments:   comments,
			CategoryID: category_id,
		}
		Transactions = append(Transactions, transaction)
	}
	fmt.Println(Transactions)
	return Transactions
}
