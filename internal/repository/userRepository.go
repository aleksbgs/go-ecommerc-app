package repository

import (
	"errors"
	"go-ecommerce-app/internal/domain"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"log"
)

type UserRepository interface {
	CreateUser(u domain.User) (domain.User, error)
	FindUserById(id int) (domain.User, error)
	FindUser(email string) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r userRepository) CreateUser(usr domain.User) (domain.User, error) {

	err := r.db.Create(&usr).Error
	if err != nil {
		log.Printf("Error while creating user: %v", err)
		return domain.User{}, errors.New("Error while creating user")
	}

	return usr, nil
}
func (r userRepository) FindUserById(id int) (domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error

	if err != nil {
		log.Printf("Error while finding user: %v", err)
		return domain.User{}, errors.New("Error while finding user")
	}

	return user, nil
}
func (r userRepository) FindUser(email string) (domain.User, error) {

	var user domain.User
	err := r.db.First(&user, "email = ?", email).Error
	if err != nil {
		log.Printf("Error while finding user: %v", err)
		return domain.User{}, errors.New("Error while finding user")
	}

	return user, nil
}
func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User
	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=", id).Updates(u).Error
	if err != nil {
		log.Printf("Error while updating user: %v", err)
		return domain.User{}, errors.New("Error while updating user")
	}
	return user, nil
}
