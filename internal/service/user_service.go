package service

import (
	"best_stock/internal/model"
	"best_stock/internal/repository"
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
	return user, nil
}
