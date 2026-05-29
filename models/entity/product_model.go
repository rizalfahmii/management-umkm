package entity

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	ID         uint
	Name       string
	PriceBuy   int
	PriceSell  int
	Stock      int
	CategoryId int
	Category   Category `json:"category"`
}
