package models

type Field struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"` // price per hour in cents
}
