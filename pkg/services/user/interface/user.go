package services

import (
	"context"

	domain "mg-inc-devops/pkg/domain"
)

type UserUseCase interface {
	FindAll(ctx context.Context) ([]domain.User, error)
	FindByID(ctx context.Context, id uint) (domain.User, error)
	FindByIDWithRole(ctx context.Context, id uint) (domain.User, error)
	Save(ctx context.Context, user domain.User) (domain.User, error)
	Delete(ctx context.Context, user domain.User) error
	FindByUserName(ctx context.Context, username string) (domain.User, error)
}
