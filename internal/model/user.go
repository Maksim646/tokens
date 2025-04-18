package model

import (
	"context"
	"errors"
	"time"
)

var (
	ErrUserNotFound  = errors.New("user not found")
	ErrInvalidUserID = errors.New("invalid user id")
	UserNotFound     = "user not found"
)

type User struct {
	ID string `db:"id"`

	Email     string    `db:"email"`
	CreatedAt time.Time `db:"created_at"`
}

type IUserRepository interface {
	CreateUserByEmail(ctx context.Context, userID string, email string) (string, error)
	GetUserByID(ctx context.Context, id string) (User, error)
}

type IUserUsecase interface {
	CreateUserByEmail(ctx context.Context, userID string, email string) (string, error)
	GetUserByID(ctx context.Context, id string) (User, error)
}
