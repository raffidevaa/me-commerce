package product

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type ProductService struct {
	repo *ProductRepository
	db   *gorm.DB
}

func NewProductService(repo *ProductRepository, db *gorm.DB) *ProductService {
	return &ProductService{
		repo: repo,
		db:   db,
	}
}

func (s *ProductService) GetProductByID(ctx context.Context, id uint) (ProductDTO, error) {
	product, err := s.repo.GetProductByID(ctx, s.db, id)
	if err != nil {
		return ProductDTO{}, errors.New("failed get a product")
	}

	productDTO := ProductDTO{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}

	return productDTO, nil
}

func (s *ProductService) GetAllProducts(ctx context.Context) (GetAllProducts, error) {
	products, err := s.repo.GetAllProducts(ctx, s.db)
	if err != nil {
		return GetAllProducts{}, errors.New("failed get all products")
	}

	var productDTOs []ProductDTO

	for _, product := range products {
		productDTO := ProductDTO{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
			Stock:       product.Stock,
		}
		productDTOs = append(productDTOs, productDTO)
	}

	res := GetAllProducts{
		Products: productDTOs,
	}

	return res, nil
}
