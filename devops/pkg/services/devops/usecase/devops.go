package usecase

import (
	"context"

	domain "mg-inc-devops/pkg/domain"
	services "mg-inc-devops/pkg/services/devops/interface"
)

type devopsUseCase struct {
	devopsRepo domain.DevopsRepository
}

func NewDevopsUseCase(repo domain.DevopsRepository) services.DevopsUseCase {
	return &devopsUseCase{
		devopsRepo: repo,
	}
}

func (c *devopsUseCase) FindAll(ctx context.Context) ([]domain.Devops, error) {
	devopss, err := c.devopsRepo.FindAll(ctx)
	return devopss, err
}

func (c *devopsUseCase) FindByID(ctx context.Context, id uint) (domain.Devops, error) {
	devops, err := c.devopsRepo.FindByID(ctx, id)
	return devops, err
}

func (c *devopsUseCase) Save(ctx context.Context, devops domain.Devops) (domain.Devops, error) {
	devops, err := c.devopsRepo.Save(ctx, devops)

	return devops, err
}

func (c *devopsUseCase) SaveAll(ctx context.Context, devops []domain.Devops) (string, error) {
	_result, err := c.devopsRepo.SaveAll(ctx, devops)

	return _result, err
}

func (c *devopsUseCase) Update(ctx context.Context, devops domain.Devops) (domain.Devops, error) {
	devops, err := c.devopsRepo.Save(ctx, devops)

	return devops, err
}

func (c *devopsUseCase) Delete(ctx context.Context, devops domain.Devops) error {
	err := c.devopsRepo.Delete(ctx, devops)

	return err
}
