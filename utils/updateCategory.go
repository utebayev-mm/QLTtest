package utils

import (
	"fmt"
	"qltTestApi/model"
)

func (r Repo) UpdateCategory(category model.Category) error {
	updateCategorySQL := `UPDATE categories SET name=$1 WHERE id=$7;`
	statement, err := r.db.Prepare(updateCategorySQL)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = statement.Exec(category.Name, category.ID)
	statement.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
