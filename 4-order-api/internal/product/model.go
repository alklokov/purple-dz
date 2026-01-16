package product

import (
	"github.com/lib/pq"
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string         `gorm:"required" json:"name" validate:"required"`
	Description string         `json:"description"`
	Price       float64        `json:"price" validate:"number"`
	Images      pq.StringArray `gorm:"type:text[]"`
}
