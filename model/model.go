package model

type Transaction struct {
	ID           int
	Name         string
	Price        int
	Date         string
	Type         bool
	Comments     string
	CategoryID   int
	CategoryName string
	Categories   []Category
}

type Category struct {
	ID   int
	Name string
}
