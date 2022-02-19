package utils

import (
	"log"
	"qltTestApi/model"
)

func (r Repo) BrowseCategories() []model.Category {
	rows, err := r.db.Query("SELECT id, name FROM categories")
	if err != nil {
		log.Println(err)
	}
	defer rows.Close()
	var Categories []model.Category
	for rows.Next() {
		var id int
		var name string
		rows.Scan(&id, &name)
		category := model.Category{
			ID:   id,
			Name: name,
		}
		Categories = append(Categories, category)
	}
	return Categories
}
