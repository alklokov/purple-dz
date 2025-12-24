package shop

import "gorm.io/gorm"

type Product struct {
	gorm.Model
	Name        string `gorm:"required"`
	Description string
	Price       float64
}
