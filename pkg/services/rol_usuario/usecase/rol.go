package usecase

import (
	"context"

	domain "mg-inc-devops/pkg/domain"
	services "mg-inc-devops/pkg/services/rol_usuario/interface"
)

type rol_UsuarioUseCase struct {
	rol_UsuarioRepo domain.Rol_UsuarioRepository
}

func NewRol_UsuarioUseCase(repo domain.Rol_UsuarioRepository) services.Rol_UsuarioUseCase {
	return &rol_UsuarioUseCase{
		rol_UsuarioRepo: repo,
	}
}

func (c *rol_UsuarioUseCase) FindAll(ctx context.Context) ([]domain.Rol_Usuario, error) {
	rol_Usuarios, err := c.rol_UsuarioRepo.FindAll(ctx)
	return rol_Usuarios, err
}

func (c *rol_UsuarioUseCase) FindByID(ctx context.Context, id uint) (domain.Rol_Usuario, error) {
	rol_Usuario, err := c.rol_UsuarioRepo.FindByID(ctx, id)
	return rol_Usuario, err
}

func (c *rol_UsuarioUseCase) FindByIDUserIDRol(ctx context.Context, idUser uint, idRol uint) ([]domain.Rol_Usuario, error) {
	rol_Usuario, err := c.rol_UsuarioRepo.FindByIDUserIDRol(ctx, idUser, idRol)
	return rol_Usuario, err
}

func (c *rol_UsuarioUseCase) Save(ctx context.Context, rol_Usuario domain.Rol_Usuario) (domain.Rol_Usuario, error) {
	rol_Usuario, err := c.rol_UsuarioRepo.Save(ctx, rol_Usuario)

	return rol_Usuario, err
}

func (c *rol_UsuarioUseCase) Update(ctx context.Context, rol_Usuario domain.Rol_Usuario) (domain.Rol_Usuario, error) {
	rol_Usuario, err := c.rol_UsuarioRepo.Save(ctx, rol_Usuario)

	return rol_Usuario, err
}

func (c *rol_UsuarioUseCase) Delete(ctx context.Context, rol_Usuario domain.Rol_Usuario) error {
	err := c.rol_UsuarioRepo.Delete(ctx, rol_Usuario)

	return err
}

func (c *rol_UsuarioUseCase) DeleteArray(ctx context.Context, rol_Usuario []domain.Rol_Usuario) error {
	err := c.rol_UsuarioRepo.DeleteArray(ctx, rol_Usuario)

	return err
}
