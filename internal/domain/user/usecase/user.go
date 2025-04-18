package usecase

import (
	"context"

	"github.com/Maksim646/tokens/internal/model"
)

type UserUsecase struct {
	userRepository model.IUserUsecase
}

func New(userRepository model.IUserRepository) model.IUserUsecase {
	return &UserUsecase{
		userRepository: userRepository,
	}
}

func (u *UserUsecase) CreateUserByEmail(ctx context.Context, userID string, email string) (string, error) {
	return u.userRepository.CreateUserByEmail(ctx, userID, email)
}

func (u *UserUsecase) GetUserByID(ctx context.Context, id string) (model.User, error) {
	return u.userRepository.GetUserByID(ctx, id)
}
