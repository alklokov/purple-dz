package product

import (
	"4-order-api/pkg/db"
	"fmt"
)

type ProductRepository struct {
	Database *db.Db
}

func NewProductRepository(database *db.Db) *ProductRepository {
	return &ProductRepository{
		Database: database,
	}
}

func (r *ProductRepository) Create(p *Product) (*Product, error) {
	res := r.Database.DB.Create(p)
	if res.Error != nil {
		return nil, res.Error
	}
	return p, nil
}

func (r *ProductRepository) GetAll() ([]Product, error) {
	var p []Product
	res := r.Database.DB.Find(&p)
	if res.Error != nil {
		return nil, res.Error
	}
	return p, nil
}

func (r *ProductRepository) GetById(id uint) (*Product, error) {
	var p []Product
	res := r.Database.DB.Find(&p, "id=?", id)
	if res.Error != nil {
		return nil, res.Error
	}
	if len(p) == 0 {
		return nil, nil
	}
	return &p[0], nil
}

func (r *ProductRepository) DeleteById(id uint) error {
	res := r.Database.DB.Delete(&Product{}, id)
	if res.Error != nil {
		return res.Error
	}
	return nil
}

func (r *ProductRepository) Update(id uint, p *Product) error {
	res := r.Database.DB.Model(&Product{}).Where("id=?", id).Updates(p)
	if res.Error != nil {
		fmt.Println("ERR Update ", res.Error)
		return res.Error
	}
	fmt.Println("RowsAffected = ", res.RowsAffected)
	return nil
}
