package utils

import (
	"fmt"
	"qltTestApi/model"
)

func (r Repo) CreateCategory(category model.Category) error {
	insertCategorySQL := `INSERT INTO categories(name) VALUES (?)`
	statement, err := r.db.Prepare(insertCategorySQL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = statement.Exec(category.Name)
	statement.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
