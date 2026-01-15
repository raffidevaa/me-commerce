package cart

import (
	"context"

	"gorm.io/gorm"
)

type CartRepository struct {
	db *gorm.DB
}

func NewCartRepository(db *gorm.DB) *CartRepository {
	return &CartRepository{
		db: db,
	}
}

func (r *CartRepository) AddItemToCart(ctx context.Context, tx *gorm.DB, cartItem CartItem, cartID uint) (CartItem, error) {
	if tx == nil {
		tx = r.db
	}

	cartItem.CartID = cartID
	if err := tx.WithContext(ctx).Create(&cartItem).Error; err != nil {
		return CartItem{}, err
	}

	return cartItem, nil
}

func (r *CartRepository) CreateCart(ctx context.Context, tx *gorm.DB, cart Cart) (Cart, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&cart).Error; err != nil {
		return Cart{}, err
	}

	return cart, nil
}

func (r *CartRepository) GetCartByUserID(ctx context.Context, tx *gorm.DB, userID uint) (Cart, error) {
	if tx == nil {
		tx = r.db
	}

	var cart Cart
	if err := tx.WithContext(ctx).Where("user_id = ?", userID).First(&cart).Error; err != nil {
		return Cart{}, err
	}

	return cart, nil
}

func (r *CartRepository) GetCartItemsByCartID(ctx context.Context, tx *gorm.DB, cartID uint) ([]CartItem, error) {
	if tx == nil {
		tx = r.db
	}

	var cartItems []CartItem
	if err := tx.WithContext(ctx).Where("cart_id = ?", cartID).Find(&cartItems).Error; err != nil {
		return nil, err
	}

	return cartItems, nil
}

func (r *CartRepository) CheckCartExistsByUserID(ctx context.Context, tx *gorm.DB, userID uint) (bool, error) {
	if tx == nil {
		tx = r.db
	}

	var count int64
	if err := tx.WithContext(ctx).Model(&Cart{}).Where("user_id = ?", userID).Count(&count).Error; err != nil {
		return false, err
	}

	return count > 0, nil
}
