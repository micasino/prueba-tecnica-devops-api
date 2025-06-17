package db

import (
	"fmt"

	"gorm.io/driver/postgres"
	//"gorm.io/driver/sqlserver"
	"gorm.io/gorm"

	config "mg-inc-devops/pkg/config"
	domain "mg-inc-devops/pkg/domain"

	"gorm.io/gorm/logger"
)

func ConnectDatabase(cfg config.Config) (*gorm.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s user=%s dbname=%s port=%s password=%s", cfg.DBHost, cfg.DBUser, cfg.DBName, cfg.DBPort, cfg.DBPassword)
	db, dbErr := gorm.Open(postgres.Open(psqlInfo), &gorm.Config{
		SkipDefaultTransaction: true,
		Logger:                 logger.Default.LogMode(logger.Info),
	})

	db.AutoMigrate(&domain.User{})
	db.AutoMigrate(&domain.Rol{})
	db.AutoMigrate(&domain.Rol_Usuario{})
	db.AutoMigrate(&domain.Devops{})

	if err := db.SetupJoinTable(&domain.User{}, "Roles", &domain.Rol_Usuario{}); err != nil {
		println(err.Error())
		panic("Failed to setup join table")
	}

	return db, dbErr
}
