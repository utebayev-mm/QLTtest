package utils

import (
	"log"
	"qltTestApi/model"
	"sort"
	"strconv"
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
		var name, date, comments, category_id, categoryName string
		var transactionType bool
		rows.Scan(&id, &name, &price, &date, &transactionType, &comments, &category_id)
		category_id_int, err := strconv.Atoi(category_id)
		if err != nil {
			log.Println(err)
		}
		categoryName = r.SelectCategory(category_id_int).Name
		transaction := model.Transaction{
			ID:           id,
			Name:         name,
			Price:        price,
			Date:         date,
			Type:         transactionType,
			Comments:     comments,
			CategoryID:   category_id_int,
			CategoryName: categoryName,
		}
		Transactions = append(Transactions, transaction)
		sort.SliceStable(Transactions, func(i, j int) bool { return Transactions[i].ID > Transactions[j].ID })

	}
	return Transactions
}

func (r Repo) BrowseByCategory(categoryID int) []model.Transaction {
	rows, err := r.db.Query("SELECT id, name, price, date, type, comments ,category_id FROM transactions WHERE category_id=?", categoryID)
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var Transactions []model.Transaction
	for rows.Next() {
		var id, price int
		var name, date, comments, category_id, categoryName string
		var transactionType bool
		rows.Scan(&id, &name, &price, &date, &transactionType, &comments, &category_id)
		category_id_int, err := strconv.Atoi(category_id)
		if err != nil {
			log.Println(err)
		}
		categoryName = r.SelectCategory(category_id_int).Name
		transaction := model.Transaction{
			ID:           id,
			Name:         name,
			Price:        price,
			Date:         date,
			Type:         transactionType,
			Comments:     comments,
			CategoryID:   category_id_int,
			CategoryName: categoryName,
		}
		Transactions = append(Transactions, transaction)
		sort.SliceStable(Transactions, func(i, j int) bool { return Transactions[i].ID > Transactions[j].ID })

	}
	return Transactions
}
