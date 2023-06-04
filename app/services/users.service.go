package service

import (
	repository "todoList_GoLang/app/repositories"
	model "todoList_GoLang/db/models"
	"todoList_GoLang/pkg"
)

type UserService interface {
	GetByCode(code string) (user *model.User, err error)
	GetById(id uint64) (*model.User, error)
	GetAll() (*[]model.User, error)
}

type userServiceImpl struct {
	repository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) UserService {
	return &userServiceImpl{repository: repository}
}

func (s *userServiceImpl) GetByCode(code string) (user *model.User, err error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(*users); i++ {
		compare := pkg.KomparePassword(code, (*users)[i].Code)
		if compare {
			return &(*users)[i], nil
		}
	}
	return nil, err
}

func (s *userServiceImpl) GetById(id uint64) (*model.User, error) {
	user, err := s.repository.GetById(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *userServiceImpl) GetAll() (*[]model.User, error) {
	users, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
