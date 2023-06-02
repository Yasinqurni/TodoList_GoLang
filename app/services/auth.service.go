package service

import (
	repository "todoList_GoLang/app/repositories"
	model "todoList_GoLang/db/models"
	"todoList_GoLang/pkg"
)

type AuthService interface {
	Login(code string) (token string, err error)
	Register(name, code string) error
}

type authServiceImpl struct {
	repository repository.UserRepository
	jwt        pkg.JWTService
}

func NewAuthRepositoryImpl(repository repository.UserRepository, jwt pkg.JWTService) AuthService {
	return &authServiceImpl{
		repository: repository,
		jwt:        jwt,
	}
}

func (s *authServiceImpl) Register(name, code string) error {
	hash, _ := pkg.HashPassword(code)
	payload := model.User{
		Name: name,
		Code: hash,
	}
	err := s.repository.Create(&payload)
	if err != nil {
		return err
	}
	return nil
}

func (s *authServiceImpl) Login(code string) (token string, err error) {
	token, err = s.jwt.GenerateToken(code)
	if err != nil {
		return "", err
	}
	return token, nil
}
