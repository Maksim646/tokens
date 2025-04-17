package model

import (
	"context"
	"time"
)

var (
	TokensAlreadyExists = "tokens already exists"
	RefreshTokenExpired = "refresh token expired"
	TokensDoNotMatch    = "tokens do not match"
)

type RefreshToken struct {
	TokenHash string `db:"token_hash"`

	ID        string    `db:"id"`
	UserID    string    `db:"user_id"`
	CreatedAt time.Time `db:"created_at"`
	ExpiredAt time.Time `db:"expired_at"`
}

type IRefreshTokenRepository interface {
	CreateRefreshToken(ctx context.Context, refreshToken RefreshToken) error
	GetRefreshTokenByUserID(ctx context.Context, userID string) (RefreshToken, error)
	DeleteRefreshTokenByID(ctx context.Context, refreshTokenID string) error
}

type IRefreshTokenUsecase interface {
	CreateRefreshToken(ctx context.Context, refreshToken RefreshToken) error
	GetRefreshTokenByUserID(ctx context.Context, userID string) (RefreshToken, error)
	DeleteRefreshTokenByID(ctx context.Context, refreshTokenID string) error
}
