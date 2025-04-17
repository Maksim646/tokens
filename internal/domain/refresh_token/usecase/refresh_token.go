package usecase

import (
	"context"

	"github.com/Maksim646/tokens/internal/model"
)

type RefreshTokenUsecase struct {
	refreshTokenRepository model.IRefreshTokenUsecase
}

func New(refreshTokenRepository model.IRefreshTokenRepository) model.IRefreshTokenUsecase {
	return &RefreshTokenUsecase{
		refreshTokenRepository: refreshTokenRepository,
	}
}

func (u *RefreshTokenUsecase) CreateRefreshToken(ctx context.Context, refreshToken model.RefreshToken) error {
	return u.refreshTokenRepository.CreateRefreshToken(ctx, refreshToken)
}

func (u *RefreshTokenUsecase) GetRefreshTokenByUserID(ctx context.Context, userID string) (model.RefreshToken, error) {
	return u.refreshTokenRepository.GetRefreshTokenByUserID(ctx, userID)
}

func (u *RefreshTokenUsecase) DeleteRefreshTokenByID(ctx context.Context, refreshTokenID string) error {
	return u.refreshTokenRepository.DeleteRefreshTokenByID(ctx, refreshTokenID)
}
