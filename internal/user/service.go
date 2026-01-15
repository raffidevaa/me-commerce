package user

import (
	"context"
	"errors"
	"time"

	"github.com/raffidevaa/me-commerce/pkg/helpers"
	"github.com/raffidevaa/me-commerce/pkg/jwtauth"
	"gorm.io/gorm"
)

type UserService struct {
	repo      *UserRepository
	tokenAuth *jwtauth.JWTAuth
	db        *gorm.DB
}

func NewUserService(repo *UserRepository, tokenAuth *jwtauth.JWTAuth, db *gorm.DB) *UserService {
	return &UserService{
		repo:      repo,
		tokenAuth: tokenAuth,
		db:        db,
	}
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

func (s *UserService) Login(ctx context.Context, req LoginRequest) (LoginResponse, error) {
	user, isExist := s.repo.FindByEmail(ctx, s.db, req.Email)
	if !isExist {
		return LoginResponse{}, errors.New("email or password is incorrect")
	}

	isValid := helpers.CheckPasswordHash(req.Password, user.Password)
	if !isValid {
		return LoginResponse{}, errors.New("email or password is incorrect")
	}

	claims := map[string]interface{}{
		"user_id": user.ID,
		"role":    user.Role,
	}

	jwtauth.SetIssuedNow(claims)
	jwtauth.SetExpiryIn(claims, 24*time.Hour)

	_, tokenString, _ := s.tokenAuth.Encode(claims)

	res := LoginResponse{
		UserID: user.ID,
		Email:  user.Email,
		Token:  tokenString,
	}

	return res, nil
}
