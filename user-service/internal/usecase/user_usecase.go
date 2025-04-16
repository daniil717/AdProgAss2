package usecase

import (
	"context"
	"user-service/internal/models"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *models.User) error
	FindByEmail(ctx context.Context, email string) (*models.User, error)
}

type UserUseCase struct {
	repo UserRepo
}

func NewUserUseCase(repo UserRepo) *UserUseCase {
	return &UserUseCase{repo: repo}
}

func (u *UserUseCase) Register(ctx context.Context, user *models.User) error {
	// Тут можно добавить проверку на существующего пользователя
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUseCase) Authenticate(ctx context.Context, email, password string) (bool, error) {
	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		return false, err
	}
	return user.Password == password, nil
}
