package product

import (
	"context"

	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) *ProductRepository {
	return &ProductRepository{
		db: db,
	}
}

func (r *ProductRepository) GetProductByID(ctx context.Context, tx *gorm.DB, id uint) (Product, error) {
	if tx == nil {
		tx = r.db
	}

	var product Product
	if err := tx.WithContext(ctx).First(&product, id).Error; err != nil {
		return Product{}, err
	}

	return product, nil
}

func (r *ProductRepository) GetAllProducts(ctx context.Context, tx *gorm.DB) ([]Product, error) {
	if tx == nil {
		tx = r.db
	}

	var products []Product
	if err := tx.WithContext(ctx).Find(&products).Error; err != nil {
		return nil, err
	}

	return products, nil
}
