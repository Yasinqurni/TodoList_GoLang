package repository

import (
	model "todoList_GoLang/db/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	//create
	Create(user *model.User) error
	//getById
	GetById(id uint64) (*model.User, error)
	//findAll
	GetAll() (*[]model.User, error)
}

type userRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepositoryImpl(db *gorm.DB) UserRepository {
	return &userRepositoryImpl{db: db}
}

func (r *userRepositoryImpl) Create(user *model.User) error {
	if err := r.db.Create(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *userRepositoryImpl) GetById(id uint64) (*model.User, error) {
	var user model.User
	err := r.db.Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepositoryImpl) GetAll() (*[]model.User, error) {
	var user []model.User
	err := r.db.Find(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
