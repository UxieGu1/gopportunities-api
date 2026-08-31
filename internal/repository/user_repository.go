package repository

import (
	"github.com/UxieGu1/gopportunities-api/internal/schemas"
	"gorm.io/gorm"
)


type UserRepository interface {
	Create(user *schemas.User) error
	GetByID(id uint) (*schemas.User, error)
	GetByEmail(email string) (*schemas.User, error)
	GetAll() ([]schemas.User, error)
	Update(user *schemas.User) error
	Delete(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *schemas.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*schemas.User, error) {
	var user schemas.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*schemas.User, error) {
	var user schemas.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAll() ([]schemas.User, error) {
	var users []schemas.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) Update(user *schemas.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&schemas.User{}, id).Error
}