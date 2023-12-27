package services

import (
	"context"

	domain "mg-inc-devops/pkg/domain"
)

type RolUseCase interface {
	FindAll(ctx context.Context) ([]domain.Rol, error)
	FindByID(ctx context.Context, id uint) (domain.Rol, error)
	Save(ctx context.Context, rol domain.Rol) (domain.Rol, error)
	Update(ctx context.Context, rol domain.Rol) (domain.Rol, error)
	Delete(ctx context.Context, rol domain.Rol) error
	DeleteRolUser(ctx context.Context, rol domain.Rol) error
}
