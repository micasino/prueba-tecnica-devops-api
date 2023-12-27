package services

import (
	"context"

	domain "mg-inc-devops/pkg/domain"
)

type DevopsUseCase interface {
	FindAll(ctx context.Context) ([]domain.Devops, error)
	FindByID(ctx context.Context, id uint) (domain.Devops, error)
	Save(ctx context.Context, devops domain.Devops) (domain.Devops, error)
	SaveAll(ctx context.Context, devops []domain.Devops) (string, error)
	Update(ctx context.Context, devops domain.Devops) (domain.Devops, error)
	Delete(ctx context.Context, devops domain.Devops) error
}
