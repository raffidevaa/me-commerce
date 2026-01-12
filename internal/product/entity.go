package product

import "time"

type Product struct {
	ID          uint   `gorm:"primaryKey"`
	Name        string `gorm:"not null"`
	Description string
	Price       int64 `gorm:"not null"`
	Stock       int   `gorm:"not null"`
	IsActive    bool  `gorm:"default:true"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
