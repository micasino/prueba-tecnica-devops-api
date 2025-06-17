package repository

import (
	"context"
	"log"

	domain "mg-inc-devops/pkg/domain"

	"gorm.io/gorm"
)

type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) domain.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) FindAll(ctx context.Context) ([]domain.User, error) {
	var users []domain.User
	//err := c.DB.Association("Roles").Find(&users).Error
	err := c.DB.Preload("Roles").Find(&users).Error

	return users, err
}

func (c *userDatabase) FindByID(ctx context.Context, id uint) (domain.User, error) {
	var user domain.User
	err := c.DB.First(&user, id).Error

	return user, err
}

func (c *userDatabase) FindByIDWithRole(ctx context.Context, id uint) (domain.User, error) {
	var user domain.User
	err := c.DB.Preload("Roles").First(&user, id).Error

	return user, err
}

func (c *userDatabase) Save(ctx context.Context, user domain.User) (domain.User, error) {
	err := c.DB.Save(&user).Error

	return user, err
}

func (c *userDatabase) Delete(ctx context.Context, user domain.User) error {
	err := c.DB.Delete(&user).Error
	log.Println("error: ", err)

	return nil
}

func (c *userDatabase) FindByUserName(ctx context.Context, username string) (domain.User, error) {
	var user domain.User
	// err := c.DB.First(&user, username).Error
	err := c.DB.Where("us_usuario = ?", username).First(&user).Error

	return user, err
}
