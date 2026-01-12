package database

import (
	"github.com/raffidevaa/me-commerce/internal/product"
	"github.com/raffidevaa/me-commerce/internal/user"
	"gorm.io/gorm"
)

func Seed(db *gorm.DB) error {
	users := []user.User{
		{Email: "admin@example.com", Password: "hashed", Role: "ADMIN"},
		{Email: "user@example.com", Password: "hashed", Role: "USER"},
	}

	for _, u := range users {
		db.FirstOrCreate(&u, user.User{Email: u.Email})
	}

	products := []product.Product{
		{Name: "Keyboard", Price: 1500000, Stock: 10},
		{Name: "Mouse", Price: 500000, Stock: 20},
	}

	for _, p := range products {
		db.FirstOrCreate(&p, product.Product{Name: p.Name})
	}

	return nil
}
