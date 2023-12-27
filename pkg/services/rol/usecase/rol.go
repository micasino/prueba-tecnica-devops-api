package usecase

import (
	"context"

	domain "mg-inc-devops/pkg/domain"
	services "mg-inc-devops/pkg/services/rol/interface"
)

type rolUseCase struct {
	rolRepo domain.RolRepository
}

func NewRolUseCase(repo domain.RolRepository) services.RolUseCase {
	return &rolUseCase{
		rolRepo: repo,
	}
}

func (c *rolUseCase) FindAll(ctx context.Context) ([]domain.Rol, error) {
	rols, err := c.rolRepo.FindAll(ctx)
	return rols, err
}

func (c *rolUseCase) FindByID(ctx context.Context, id uint) (domain.Rol, error) {
	rol, err := c.rolRepo.FindByID(ctx, id)
	return rol, err
}

func (c *rolUseCase) Save(ctx context.Context, rol domain.Rol) (domain.Rol, error) {
	rol, err := c.rolRepo.Save(ctx, rol)

	return rol, err
}

func (c *rolUseCase) Update(ctx context.Context, rol domain.Rol) (domain.Rol, error) {
	rol, err := c.rolRepo.Save(ctx, rol)

	return rol, err
}

func (c *rolUseCase) Delete(ctx context.Context, rol domain.Rol) error {
	err := c.rolRepo.Delete(ctx, rol)

	return err
}

func (c *rolUseCase) DeleteRolUser(ctx context.Context, rol domain.Rol) error {
	err := c.rolRepo.Delete(ctx, rol)

	return err
}
