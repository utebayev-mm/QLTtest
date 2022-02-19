package model

type Transaction struct {
	ID         int
	Name       string
	Price      int
	Date       string
	Type       string
	Comments   string
	CategoryID string
}

type Category struct {
	ID   int
	Name string
}
