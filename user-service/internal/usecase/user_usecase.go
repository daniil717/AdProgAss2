package usecase

import (
	"context"
	"log"
	"user_service/internal/models"
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
	log.Printf("[Register] received user: %+v\n", user)

	err := u.repo.CreateUser(ctx, user)
	if err != nil {
		log.Printf("[Register ERROR] failed to create user: %v\n", err)
		return err
	}
	log.Println("[Register] user created successfully")
	return nil
}

func (u *UserUseCase) Authenticate(ctx context.Context, email, password string) (bool, error) {
	log.Printf("[Authenticate] email: %s\n", email)

	user, err := u.repo.FindByEmail(ctx, email)
	if err != nil {
		log.Printf("[Authenticate ERROR] %v\n", err)
		return false, err
	}
	match := user.Password == password
	log.Printf("[Authenticate] match: %v\n", match)
	return match, nil
}
