package services

import (
	"context"

	domain "mg-inc-devops/pkg/domain"
)

type Rol_UsuarioUseCase interface {
	FindAll(ctx context.Context) ([]domain.Rol_Usuario, error)
	FindByID(ctx context.Context, id uint) (domain.Rol_Usuario, error)
	FindByIDUserIDRol(ctx context.Context, idUser uint, idRol uint) ([]domain.Rol_Usuario, error)
	Save(ctx context.Context, rol_Usuario domain.Rol_Usuario) (domain.Rol_Usuario, error)
	Update(ctx context.Context, rol_Usuario domain.Rol_Usuario) (domain.Rol_Usuario, error)
	Delete(ctx context.Context, rol_Usuario domain.Rol_Usuario) error
	DeleteArray(ctx context.Context, rol_Usuario []domain.Rol_Usuario) error
}
