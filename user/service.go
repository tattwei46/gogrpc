package user

import (
	"context"
	"fmt"
)

type Service interface {
	ListUsers(ctx context.Context) []User
}

type userService struct {
}

func NewService() (Service, error) {
	return &userService{}, nil
}

func (u userService) ListUsers(ctx context.Context) []User {
	var users []User
	for i := 0; i < 10; i++ {
		user := User{
			ID:       fmt.Sprintf("%d", i),
			UserID:   fmt.Sprintf("UserID %d", i),
			Email:    fmt.Sprintf("Email %d", i),
			Password: fmt.Sprintf("Password %d", i),
		}
		users = append(users, user)
	}

	return users
}
