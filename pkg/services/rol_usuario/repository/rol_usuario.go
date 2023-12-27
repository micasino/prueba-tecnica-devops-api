package repository

import (
	"context"

	domain "mg-inc-devops/pkg/domain"

	"gorm.io/gorm"
)

type rol_UsuarioDatabase struct {
	DB *gorm.DB
}

func NewRol_UsuarioRepository(DB *gorm.DB) domain.Rol_UsuarioRepository {
	return &rol_UsuarioDatabase{DB}
}

func (c *rol_UsuarioDatabase) FindAll(ctx context.Context) ([]domain.Rol_Usuario, error) {
	var rol_Usuarios []domain.Rol_Usuario
	err := c.DB.Where("rol_Usuario_eliminado = ? and rol_Usuario_esactivo = ?", "false", "true").Find(&rol_Usuarios).Error

	return rol_Usuarios, err
}

func (c *rol_UsuarioDatabase) FindByID(ctx context.Context, id uint) (domain.Rol_Usuario, error) {
	var rol_Usuario domain.Rol_Usuario
	err := c.DB.Where("rol_Usuario_eliminado = ?", "false").First(&rol_Usuario, id).Error

	return rol_Usuario, err
}

func (c *rol_UsuarioDatabase) FindByIDUserIDRol(ctx context.Context, idUser uint, idRol uint) ([]domain.Rol_Usuario, error) {
	var rol_Usuarios []domain.Rol_Usuario
	err := c.DB.Where("user_id = ? and rol_id = ?", idUser, idRol).Find(&rol_Usuarios).Error

	return rol_Usuarios, err
}

func (c *rol_UsuarioDatabase) Save(ctx context.Context, rol_Usuario domain.Rol_Usuario) (domain.Rol_Usuario, error) {
	err := c.DB.Save(&rol_Usuario).Error

	return rol_Usuario, err
}

func (c *rol_UsuarioDatabase) Update(ctx context.Context, rol_Usuario domain.Rol_Usuario) (domain.Rol_Usuario, error) {
	err := c.DB.Save(&rol_Usuario).Error

	return rol_Usuario, err
}

func (c *rol_UsuarioDatabase) Delete(ctx context.Context, rol_Usuario domain.Rol_Usuario) error {
	err := c.DB.Delete(&rol_Usuario).Error

	return err
}

func (c *rol_UsuarioDatabase) DeleteArray(ctx context.Context, rol_Usuario []domain.Rol_Usuario) error {
	err := c.DB.Delete(&rol_Usuario).Error

	return err
}
