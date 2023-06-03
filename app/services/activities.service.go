package service

import (
	repository "todoList_GoLang/app/repositories"
	model "todoList_GoLang/db/models"
	"todoList_GoLang/dto"
)

type ActivityService interface {
	//create
	Create(activity *dto.ActivityRequest, userId, titleId uint) error
	//bulk create
	CreateBulk(activities *[]dto.ActivityRequest, userId, titleId uint) error
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

type activityServiceImpl struct {
	activity repository.ActivityRepository
}

func NewActivityServiceImpl(activity repository.ActivityRepository) ActivityService {
	return &activityServiceImpl{activity: activity}
}

func (s *activityServiceImpl) Create(activity *dto.ActivityRequest, userId, titleId uint) error {
	var data = model.Activity{
		UserId:  userId,
		TitleId: titleId,
		List:    activity.List,
		Done:    false,
	}
	err := s.activity.Create(&data)
	if err != nil {
		return err
	}

	return nil
}

func (s *activityServiceImpl) CreateBulk(activities *[]dto.ActivityRequest, userId, titleId uint) error {

	for i := 0; i < len(*activities); i++ {
		var data = model.Activity{
			UserId:  userId,
			TitleId: titleId,
			List:    (*activities)[i].List,
			Done:    false,
		}
		err := s.activity.Create(&data)
		if err != nil {
			return err
		}
	}
	return nil

}

func (s *activityServiceImpl) Update(activity string, id uint) error {
	err := s.activity.Update(activity, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *activityServiceImpl) GetById(id, userId uint) (*model.Activity, error) {
	data, err := s.activity.GetById(id, userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *activityServiceImpl) GetAll(userId, titleId uint) (*[]model.Activity, error) {
	data, err := s.activity.GetAll(userId, titleId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s *activityServiceImpl) UpdateStatus(id uint) error {
	err := s.activity.UpdateStatus(id)
	if err != nil {
		return err
	}
	return nil
}

func (s *activityServiceImpl) Delete(id uint) error {
	err := s.activity.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
