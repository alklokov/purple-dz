package shop

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string `gorm:"required"`
	Description string
	Price       float64
	Images      pq.StringArray
}
