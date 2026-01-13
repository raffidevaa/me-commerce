package user

import (
	"context"
	"errors"

	"github.com/raffidevaa/me-commerce/pkg/helpers"
	"gorm.io/gorm"
)

type UserService struct {
	repo *UserRepository
	db   *gorm.DB
}

func NewUserService(repo *UserRepository, db *gorm.DB) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(ctx context.Context, u User) (User, error) {
	_, isExist := s.repo.FindByEmail(ctx, s.db, u.Email)
	if isExist {
		return User{}, errors.New("This email already registered")
	}

	hashedPw, err := helpers.HashPassword(u.Password)
	if err != nil {
		return User{}, errors.New("failed hash password")
	}

	user := User{
		Email:    u.Email,
		Password: hashedPw,
	}

	insertedUser, err := s.repo.Save(ctx, s.db, user)
	if err != nil {
		return User{}, errors.New("failed register user")
	}

	return insertedUser, nil

}
