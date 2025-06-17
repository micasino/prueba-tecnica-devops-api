package repository

import (
	"context"

	domain "mg-inc-devops/pkg/domain"

	"gorm.io/gorm"
)

type devopsDatabase struct {
	DB *gorm.DB
}

func NewDevopsRepository(DB *gorm.DB) domain.DevopsRepository {
	return &devopsDatabase{DB}
}

func (c *devopsDatabase) FindAll(ctx context.Context) ([]domain.Devops, error) {
	var devopss []domain.Devops
	err := c.DB.Where("deleted = ?", "false").Find(&devopss).Error

	return devopss, err
}

func (c *devopsDatabase) FindByID(ctx context.Context, id uint) (domain.Devops, error) {
	var devops domain.Devops
	err := c.DB.Where("deleted = ?", "false").First(&devops, id).Error

	return devops, err
}

func (c *devopsDatabase) Save(ctx context.Context, devops domain.Devops) (domain.Devops, error) {
	err := c.DB.Save(&devops).Error

	return devops, err
}

func (c *devopsDatabase) SaveAll(ctx context.Context, devops []domain.Devops) (string, error) {
	err := c.DB.Save(&devops).Error

	return "save all ok", err
}
func (c *devopsDatabase) Update(ctx context.Context, devops domain.Devops) (domain.Devops, error) {
	err := c.DB.Save(&devops).Error

	return devops, err
}

func (c *devopsDatabase) Delete(ctx context.Context, devops domain.Devops) error {
	err := c.DB.Delete(&devops).Error

	return err
}
