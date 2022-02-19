package utils

import (
	"fmt"
	"qltTestApi/model"
)

func (r Repo) Update(transaction model.Transaction) error {
	fmt.Println(transaction)
	updateTransactionSQL := `UPDATE transactions SET name=$1, price=$2, date=$3, type=$4, comments=$5, category_id=$6 WHERE id=$7;`
	statement, err := r.db.Prepare(updateTransactionSQL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = statement.Exec(transaction.Name, transaction.Price, transaction.Date, transaction.Type, transaction.Comments, transaction.CategoryID, transaction.ID)
	statement.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
