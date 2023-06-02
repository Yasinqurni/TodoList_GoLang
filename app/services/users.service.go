package service

import (
	repository "todoList_GoLang/app/repositories"
	model "todoList_GoLang/db/models"
	"todoList_GoLang/pkg"
)

type UserService interface {
	GetByCode(code string, id uint64) (user *model.User, err error)
}

type userServiceImpl struct {
	repository repository.UserRepository
}

func NewUserRepositoryImpl(repository repository.UserRepository) UserService {
	return &userServiceImpl{repository: repository}
}

func (s *userServiceImpl) GetByCode(code string, id uint64) (user *model.User, err error) {
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
