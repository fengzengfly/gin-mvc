package service

import (
	"gin-mvc/internal/model"
	"gin-mvc/internal/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) Register(user *model.User) error {
	// 注册逻辑
	return s.repo.Create(user)
}

func (s *UserService) Login(username, password string) (*model.User, error) {
	// 登录验证逻辑
	user, err := s.repo.FindByUsername(username)
	if err != nil {
		return nil, err
	}
	// 密码验证
	if user.Password != password {
		return nil, nil
	}

	return user, nil
}
