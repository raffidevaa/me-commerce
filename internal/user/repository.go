package user

import (
	"context"

	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Save(ctx context.Context, tx *gorm.DB, u User) (User, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&u).Error; err != nil {
		return User{}, err
	}
	return u, nil
}

func (r *UserRepository) FindByEmail(ctx context.Context, tx *gorm.DB, email string) (User, bool) {
	if tx == nil {
		tx = r.db
	}

	var u User
	if err := tx.WithContext(ctx).Where("email = ?", email).First(&u).Error; err != nil {
		return User{}, false
	}

	return u, true
}
