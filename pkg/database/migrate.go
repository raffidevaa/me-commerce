package database

import (
	"github.com/raffidevaa/me-commerce/internal/cart"
	"github.com/raffidevaa/me-commerce/internal/notification"
	"github.com/raffidevaa/me-commerce/internal/order"
	"github.com/raffidevaa/me-commerce/internal/payment"
	"github.com/raffidevaa/me-commerce/internal/product"
	"github.com/raffidevaa/me-commerce/internal/user"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user.User{},
		&product.Product{},
		&cart.Cart{},
		&cart.CartItem{},
		&order.Order{},
		&order.OrderItem{},
		&payment.Payment{},
		&notification.Notification{},
	)
}
