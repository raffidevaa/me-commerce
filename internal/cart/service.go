package cart

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

type CartService struct {
	repo *CartRepository
	db   *gorm.DB
}

func NewCartService(repo *CartRepository, db *gorm.DB) *CartService {
	return &CartService{
		repo: repo,
		db:   db,
	}
}

func (s *CartService) AddItemToCart(ctx context.Context, req AddItemToCartRequest, userID uint) (CartItem, error) {
	isExist, _ := s.CheckCartExistsByUserID(ctx, userID)

	if isExist == false {
		newCart := Cart{
			UserID: userID,
		}

		_, err := s.repo.CreateCart(ctx, s.db, newCart)
		if err != nil {
			return CartItem{}, errors.New("failed to create cart for user")
		}
	}

	cart, err := s.repo.GetCartByUserID(ctx, s.db, userID)
	if err != nil {
		return CartItem{}, errors.New("failed to get cart for user")
	}

	cartItem := CartItem{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
	}

	insertedCartItem, err := s.repo.AddItemToCart(ctx, s.db, cartItem, cart.ID)
	if err != nil {
		return CartItem{}, errors.New("failed to add item to cart")
	}

	return insertedCartItem, nil
}

func (s *CartService) GetCartByUserID(ctx context.Context, userID uint) (GetCartsResponse, error) {
	cart, err := s.repo.GetCartByUserID(ctx, s.db, userID)
	if err != nil {
		return GetCartsResponse{}, err
	}

	cartItems, err := s.repo.GetCartItemsByCartID(ctx, s.db, cart.ID)
	if err != nil {
		return GetCartsResponse{}, err
	}

	return GetCartsResponse{CartItems: cartItems}, nil
}

func (s *CartService) CheckCartExistsByUserID(ctx context.Context, userID uint) (bool, error) {
	return s.repo.CheckCartExistsByUserID(ctx, s.db, userID)
}
