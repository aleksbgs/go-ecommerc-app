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
	FindUserById(id uint) (domain.User, error)
	FindUser(email string) (domain.User, error)
	UpdateUser(id uint, u domain.User) (domain.User, error)
	CreateBankAccount(e domain.BankAccount) error

	// Cart
	FindCartItems(uId uint) ([]domain.Cart, error)
	FindCartItem(uId uint, pId uint) (domain.Cart, error)
	CreateCart(c domain.Cart) error
	UpdateCart(c domain.Cart) error
	DeleteCartById(id uint) error
	DeleteCartItems(uId uint) error

	// Profile
	CreateProfile(u domain.Address) error
	UpdateProfile(u domain.Address) error
}

type userRepository struct {
	db *gorm.DB
}

func (r userRepository) CreateProfile(u domain.Address) error {
	err := r.db.Create(&u).Error
	if err != nil {
		log.Printf("Error while creating profile: %v", err)
		return err
	}
	return nil
}

func (r userRepository) UpdateProfile(u domain.Address) error {
	err := r.db.Where("user_id = ?", u.UserId).Updates(u).Error
	if err != nil {
		log.Printf("Error while updating profile: %v", err)
		return err
	}
	return nil
}

func (r userRepository) FindCartItems(uId uint) ([]domain.Cart, error) {
	var carts []domain.Cart
	err := r.db.Where("user_id = ?", uId).Find(&carts).Error
	if err != nil {
		return nil, err
	}
	return carts, nil
}

func (r userRepository) FindCartItem(uId uint, pId uint) (domain.Cart, error) {
	cartItem := domain.Cart{}
	err := r.db.Where("user_id = ? AND product_id=?", uId, pId).First(&cartItem).Error
	if err != nil {
		return cartItem, err
	}
	return cartItem, nil
}

func (r userRepository) CreateCart(c domain.Cart) error {
	return r.db.Create(&c).Error
}

func (r userRepository) UpdateCart(c domain.Cart) error {
	var cart domain.Cart
	err := r.db.Model(&cart).Clauses(clause.Returning{}).Where("id = ?", c.ID).Updates(c).Error
	return err
}

func (r userRepository) DeleteCartById(id uint) error {
	err := r.db.Delete(&domain.Cart{}, id).Error
	return err
}

func (r userRepository) DeleteCartItems(uId uint) error {
	err := r.db.Where("user_id = ?", uId).Delete(&domain.Cart{}).Error
	return err
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
func (r userRepository) FindUserById(id uint) (domain.User, error) {
	var user domain.User
	err := r.db.Preload("Address").First(&user, id).Error

	if err != nil {
		log.Printf("Error while finding user: %v", err)
		return domain.User{}, errors.New("Error while finding user")
	}

	return user, nil
}
func (r userRepository) FindUser(email string) (domain.User, error) {

	var user domain.User

	err := r.db.Preload("Address").First(&user, "email = ?", email).Error

	if err != nil {
		log.Printf("Error while finding user: %v", err)
		return domain.User{}, errors.New("Error while finding user")
	}

	return user, nil
}
func (r userRepository) UpdateUser(id uint, u domain.User) (domain.User, error) {
	var user domain.User
	err := r.db.Model(&user).Clauses(clause.Returning{}).Where("id=?", id).Updates(u).Error
	if err != nil {
		log.Printf("Error while updating user: %v", err)
		return domain.User{}, errors.New("Error while updating user")
	}
	return user, nil
}
func (r userRepository) CreateBankAccount(e domain.BankAccount) error {
	err := r.db.Create(&e).Error
	if err != nil {
		return errors.New("Error while creating bank account")
	}
	return err
}
