package utils

import (
	"log"
	"qltTestApi/model"
)

func (r Repo) SelectCategory(IDToEdit int) model.Category {
	var id int
	var name string
	err := r.db.QueryRow("SELECT id, name FROM categories WHERE id = ?", IDToEdit).Scan(&id, &name)
	if err != nil {
		log.Println(err)
	}
	Category := model.Category{
		ID:   id,
		Name: name,
	}
	return Category
}

func (r Repo) SelectCategorybyName(name string) model.Category {
	var id int
	err := r.db.QueryRow("SELECT id, name FROM categories WHERE name = ?", name).Scan(&id, &name)
	if err != nil {
		log.Println(err)
	}
	Category := model.Category{
		ID:   id,
		Name: name,
	}
	return Category
}
