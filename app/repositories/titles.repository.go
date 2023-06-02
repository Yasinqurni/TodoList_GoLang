package repository

import (
	model "todoList_GoLang/db/models"

	"gorm.io/gorm"
)

type TitleRepository interface {
	//create
	Create(title *model.Title) error
	//getall
	GetAll(userId uint) (*[]model.Title, error)
	//getbyid
	GetById(id, userId uint) (*model.Title, error)
	//update
	Update(title string, id uint) error
	//delete
	Delete(id uint) error
}

type titleRepositoryImpl struct {
	db *gorm.DB
}

func NewTitleRepositoryImpl(db *gorm.DB) TitleRepository {
	return &titleRepositoryImpl{db: db}
}

func (r *titleRepositoryImpl) Create(title *model.Title) error {
	if err := r.db.Create(title).Error; err != nil {
		return err
	}

	return nil
}

func (r *titleRepositoryImpl) GetAll(userId uint) (*[]model.Title, error) {
	var title []model.Title
	err := r.db.Find(&title).Where("user_id", userId).Error
	if err != nil {
		return nil, err
	}
	return &title, nil
}

func (r *titleRepositoryImpl) GetById(id, userId uint) (*model.Title, error) {
	var title model.Title
	err := r.db.Where("id = ?", id).Where("user_id = ?", userId).Take(&title).Error
	if err != nil {
		return nil, err
	}
	return &title, nil
}

func (r *titleRepositoryImpl) Update(title string, id uint) error {
	err := r.db.Where("id = ?", id).Update("title", title).Error
	if err != nil {
		return err
	}
	return nil

}

func (r *titleRepositoryImpl) Delete(id uint) error {
	var title model.Title
	err := r.db.Where("id = ?", id).Delete(&title).Error
	if err != nil {
		return err
	}
	return nil
}
