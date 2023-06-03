package repository

import (
	model "todoList_GoLang/db/models"

	"gorm.io/gorm"
)

type ActivityRepository interface {
	//create
	Create(activity *model.Activity) error
	//bulk create
	// CreateBulk(activity *[]model.Activity) error
	//update
	Update(activity string, id uint) error
	//getbyid
	GetById(id, userId uint) (*model.Activity, error)
	//getall
	GetAll(userId, titleId uint) (*[]model.Activity, error)
	//updatestatus
	UpdateStatus(id uint) error
	//delete
	Delete(id uint) error
}

type activityRepositoryImpl struct {
	db *gorm.DB
}

func NewActivityRepositoryImpl(db *gorm.DB) ActivityRepository {
	return &activityRepositoryImpl{db: db}
}

func (r *activityRepositoryImpl) Create(activity *model.Activity) error {
	if err := r.db.Create(activity).Error; err != nil {
		return err
	}

	return nil
}

// func (r *activityRepositoryImpl) CreateBulk(activity *[]model.Activity) error {
// 	if err := r.db.Create(activity).Error; err != nil {
// 		return err
// 	}

// 	return nil
// }

func (r *activityRepositoryImpl) Update(activity string, id uint) error {
	err := r.db.Where("id = ?", id).Update("list", activity).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *activityRepositoryImpl) GetById(id, userId uint) (*model.Activity, error) {
	var activity model.Activity
	err := r.db.Where("id = ?", id).Where("user_id = ?", userId).Take(&activity).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *activityRepositoryImpl) GetAll(userId, titleId uint) (*[]model.Activity, error) {
	var activity []model.Activity
	err := r.db.Find(&activity).Where("user_id", userId).Where("title_id", titleId).Error
	if err != nil {
		return nil, err
	}
	return &activity, nil
}

func (r *activityRepositoryImpl) UpdateStatus(id uint) error {
	err := r.db.Where("id = ?", id).Update("done", true).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *activityRepositoryImpl) Delete(id uint) error {
	var title model.Title
	err := r.db.Where("id = ?", id).Delete(&title).Error
	if err != nil {
		return err
	}
	return nil
}
