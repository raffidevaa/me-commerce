package order

import (
	"context"

	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{
		db: db,
	}
}

func (r *OrderRepository) CreateOrder(ctx context.Context, order Order) (Order, error) {
	tx := r.db.WithContext(ctx)

	if err := tx.Create(&order).Error; err != nil {
		return Order{}, err
	}

	return order, nil
}

func (r *OrderRepository) CreateOrderItems(ctx context.Context, orderItem OrderItem) error {
	tx := r.db.WithContext(ctx)

	if err := tx.Create(&orderItem).Error; err != nil {
		return err
	}

	return nil
}

func (r *OrderRepository) GetOrdersByUserID(ctx context.Context, userID uint) ([]Order, error) {
	tx := r.db.WithContext(ctx)

	var orders []Order
	if err := tx.Where("user_id = ?", userID).Find(&orders).Error; err != nil {
		return nil, err
	}

	return orders, nil
}

func (r *OrderRepository) GetOrderByID(ctx context.Context, orderID uint) (Order, error) {
	tx := r.db.WithContext(ctx)

	var order Order
	if err := tx.Where("id = ?", orderID).First("id", orderID).Error; err != nil {
		return Order{}, err
	}

	return order, nil
}

func (r *OrderRepository) Payment(ctx context.Context, orderID uint) (Order, error) {
	tx := r.db.WithContext(ctx)
	
	var order Order

	if err := tx.Where("id = ?", orderID).Update("status", StatusPaid).Error; err != nil {
		return Order{}, err
	}

	return order, nil
}



