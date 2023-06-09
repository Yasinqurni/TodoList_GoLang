package service

import (
	repository "todoList_GoLang/app/repositories"
	model "todoList_GoLang/db/models"
	"todoList_GoLang/dto"
)

type TitleService interface {
	//create
	Create(title *dto.TitleRequest, userId uint) error
	//getall
	GetAll(userId uint) (*[]model.Title, error)
	//getbyid
	GetById(id, userId uint) (*model.Title, error)
	//update
	Update(title dto.TitleRequest, id uint) error
	//delete
	Delete(id uint) error
}

type titleServiceImpl struct {
	repository repository.TitleRepository
}

func NewTitleServiceImpl(repository repository.TitleRepository) TitleService {
	return &titleServiceImpl{repository: repository}
}

func (s titleServiceImpl) Create(title *dto.TitleRequest, userId uint) error {
	data := model.Title{
		UserId: userId,
		Title:  title.Title,
	}
	err := s.repository.Create(&data)
	if err != nil {
		return err
	}
	return nil
}

func (s titleServiceImpl) GetAll(userId uint) (*[]model.Title, error) {
	data, err := s.repository.GetAll(userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s titleServiceImpl) GetById(id, userId uint) (*model.Title, error) {
	data, err := s.repository.GetById(id, userId)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (s titleServiceImpl) Update(title dto.TitleRequest, id uint) error {

	var data = title.Title
	err := s.repository.Update(data, id)
	if err != nil {
		return err
	}
	return nil
}

func (s titleServiceImpl) Delete(id uint) error {
	err := s.repository.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
