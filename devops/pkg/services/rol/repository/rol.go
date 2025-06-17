package repository

import (
	"context"

	domain "mg-inc-devops/pkg/domain"

	"gorm.io/gorm"
)

type rolDatabase struct {
	DB *gorm.DB
}

func NewRolRepository(DB *gorm.DB) domain.RolRepository {
	return &rolDatabase{DB}
}

func (c *rolDatabase) FindAll(ctx context.Context) ([]domain.Rol, error) {
	var rols []domain.Rol
	err := c.DB.Where("rol_eliminado = ? and rol_esactivo = ?", "false", "true").Find(&rols).Error

	return rols, err
}

func (c *rolDatabase) FindByID(ctx context.Context, id uint) (domain.Rol, error) {
	var rol domain.Rol
	err := c.DB.Where("rol_eliminado = ?", "false").First(&rol, id).Error

	return rol, err
}

func (c *rolDatabase) Save(ctx context.Context, rol domain.Rol) (domain.Rol, error) {
	err := c.DB.Save(&rol).Error

	return rol, err
}

func (c *rolDatabase) Update(ctx context.Context, rol domain.Rol) (domain.Rol, error) {
	err := c.DB.Save(&rol).Error

	return rol, err
}

func (c *rolDatabase) Delete(ctx context.Context, rol domain.Rol) error {
	err := c.DB.Delete(&rol).Error

	return err
}

func (c *rolDatabase) DeleteRolUser(ctx context.Context, rol domain.Rol) error {
	//err := c.DB.Delete(&rol).Error
	err := c.DB.Select("Roles").Delete(&rol).Error

	return err
}
