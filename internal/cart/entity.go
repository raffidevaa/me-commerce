package cart

import "time"

type CartItem struct {
	ID        uint `gorm:"primaryKey"`
	CartID    uint `gorm:"index"`
	ProductID uint
	Quantity  int
}

type Cart struct {
	ID        uint `gorm:"primaryKey"`
	UserID    uint `gorm:"uniqueIndex"`
	UpdatedAt time.Time
}
