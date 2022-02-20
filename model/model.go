package model

type Transaction struct {
	ID           int
	Name         string
	Price        int
	Date         string
	Type         bool
	Comments     string
	CategoryID   string
	CategoryName string
}

type Category struct {
	ID   int
	Name string
}
