package usecase

import (
	"context"

	domain "mg-inc-devops/pkg/domain"
	services "mg-inc-devops/pkg/services/user/interface"
)

type userUseCase struct {
	userRepo domain.UserRepository
}

func NewUserUseCase(repo domain.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) FindAll(ctx context.Context) ([]domain.User, error) {
	users, err := c.userRepo.FindAll(ctx)
	return users, err
}

func (c *userUseCase) FindByID(ctx context.Context, id uint) (domain.User, error) {
	user, err := c.userRepo.FindByID(ctx, id)
	return user, err
}

func (c *userUseCase) FindByIDWithRole(ctx context.Context, id uint) (domain.User, error) {
	user, err := c.userRepo.FindByIDWithRole(ctx, id)
	return user, err
}

func (c *userUseCase) Save(ctx context.Context, user domain.User) (domain.User, error) {
	user, err := c.userRepo.Save(ctx, user)

	return user, err
}

func (c *userUseCase) Delete(ctx context.Context, user domain.User) error {
	err := c.userRepo.Delete(ctx, user)

	return err
}

func (c *userUseCase) FindByUserName(ctx context.Context, username string) (domain.User, error) {
	user, err := c.userRepo.FindByUserName(ctx, username)

	return user, err
}
