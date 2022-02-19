package utils

import (
	"database/sql"
	"fmt"
	"qltTestApi/model"
)

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{db: db}
}

// Create - inserts the records into transactions table
func (r Repo) Create(transaction model.Transaction) error {
	insertTransactionSQL := `INSERT INTO transactions(name, price, date, type, comments ,category_id) VALUES (?, ?, ?, ?, ?, ?)`
	statement, err := r.db.Prepare(insertTransactionSQL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = statement.Exec(transaction.Name, transaction.Price, transaction.Date, transaction.Type, transaction.Comments, transaction.CategoryID)
	statement.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
