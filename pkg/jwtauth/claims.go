package jwtauth

import (
	"context"
)

type Claims struct {
	UserID uint
	Role   string
}

func GetClaims(ctx context.Context) (*Claims, bool) {
	_, claims, err := FromContext(ctx)
	if err != nil {
		return nil, false
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return nil, false
	}

	role, ok := claims["role"].(string)
	if !ok {
		return nil, false
	}

	return &Claims{
		UserID: uint(userID),
		Role:   role,
	}, true
}
