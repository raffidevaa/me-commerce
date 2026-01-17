package order

import (
	"context"
	"errors"

	"github.com/raffidevaa/me-commerce/internal/cart"
	"github.com/raffidevaa/me-commerce/internal/product"
	"gorm.io/gorm"
)

type OrderService struct {
	orderRepository   *OrderRepository
	cartRepository    *cart.CartRepository
	productRepository *product.ProductRepository
	db                *gorm.DB
}

func NewOrderService(orderRepository *OrderRepository, cartRepository *cart.CartRepository, productRepository *product.ProductRepository, db *gorm.DB) *OrderService {
	return &OrderService{
		orderRepository:   orderRepository,
		cartRepository:    cartRepository,
		productRepository: productRepository,
		db:                db,
	}
}

func (s *OrderService) CreateOrder(ctx context.Context, req CreateOrderRequest, userID uint) (CreateOrderResponse, error) {
	tx := s.db.WithContext(ctx)

	cartItem, err := s.cartRepository.GetCartItemByCartItemID(ctx, tx, req.CartItemID)
	if err != nil {
		return CreateOrderResponse{}, errors.New("failed get cart item")
	}

	product, err := s.productRepository.GetProductByID(ctx, tx, cartItem.ProductID)
	if err != nil {
		return CreateOrderResponse{}, errors.New("failed get product")
	}

	order := Order{
		UserID: userID,
		Total:  product.Price * int64(cartItem.Quantity),
		Status: string(StatusPending),
	}

	createdOrder, err := s.orderRepository.CreateOrder(ctx, order)
	if err != nil {
		return CreateOrderResponse{}, errors.New("failed create order")
	}

	orderItem := OrderItem{
		OrderID:   createdOrder.ID,
		ProductID: product.ID,
		Quantity:  cartItem.Quantity,
		Price:     product.Price,
	}

	err = s.orderRepository.CreateOrderItems(ctx, orderItem)
	if err != nil {
		return CreateOrderResponse{}, errors.New("failed create order items")
	}

	// update cart item as already purchased
	err = s.cartRepository.UpdateCartItemIsAlreadyPurchased(ctx, tx, cartItem.ID, true)
	if err != nil {
		return CreateOrderResponse{}, errors.New("failed update cart item as already purchased")
	}

	res := CreateOrderResponse{
		UserID:    createdOrder.UserID,
		OrderID:   createdOrder.ID,
		ProductID: orderItem.ProductID,
		UnitPrice: orderItem.Price,
		Quantity:  orderItem.Quantity,
		Total:     createdOrder.Total,
		Status:    createdOrder.Status,
	}

	return res, nil
}

func (s *OrderService) GetOrdersByUserID(ctx context.Context, userID uint) ([]Order, error) {
	orders, err := s.orderRepository.GetOrdersByUserID(ctx, userID)
	if err != nil {
		return nil, errors.New("failed get orders by user id")
	}

	return orders, nil
}
