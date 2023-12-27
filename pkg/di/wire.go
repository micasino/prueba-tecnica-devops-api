//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"
	config "mg-inc-devops/pkg/config"
	db "mg-inc-devops/pkg/db"
	http "mg-inc-devops/pkg/http"
	handler "mg-inc-devops/pkg/http/handler"
	devopsrepository "mg-inc-devops/pkg/services/devops/repository"
	devopsusecase "mg-inc-devops/pkg/services/devops/usecase"
	rolrepository "mg-inc-devops/pkg/services/rol/repository"
	rolusecase "mg-inc-devops/pkg/services/rol/usecase"
	rol_usuariorepository "mg-inc-devops/pkg/services/rol_usuario/repository"
	rol_usuariousecase "mg-inc-devops/pkg/services/rol_usuario/usecase"
	userrepository "mg-inc-devops/pkg/services/user/repository"
	userusecase "mg-inc-devops/pkg/services/user/usecase"
)

func InitializeAPI(cfg config.Config) (*http.ServerHTTP, error) {
	wire.Build(db.ConnectDatabase,
		userrepository.NewUserRepository,
		userusecase.NewUserUseCase,
		handler.NewUserHandler,
		rolrepository.NewRolRepository,
		rolusecase.NewRolUseCase,
		handler.NewRolHandler,
		rol_usuariorepository.NewRol_UsuarioRepository,
		rol_usuariousecase.NewRol_UsuarioUseCase,
		devopsrepository.NewDevopsRepository,
		devopsusecase.NewDevopsUseCase,
		handler.NewDevopsHandler,
		http.NewServerHTTP)

	return &http.ServerHTTP{}, nil
}
