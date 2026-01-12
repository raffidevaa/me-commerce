package order

import "time"

type Status string

const (
	StatusPending   Status = "PENDING"
	StatusPaid      Status = "PAID"
	StatusCancelled Status = "CANCELLED"
)

type Order struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    uint   `gorm:"not null"`
	Total     int64  `gorm:"not null"`
	Status    string `gorm:"type:varchar(20)"`
	Items     []OrderItem
	CreatedAt time.Time
	UpdatedAt time.Time
}

type OrderItem struct {
	ID        uint `gorm:"primaryKey"`
	OrderID   uint
	ProductID uint
	Quantity  int
	Price     int64
}
