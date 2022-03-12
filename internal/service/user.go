package service

import (
	"architecture/internal/core"
	"context"
)

type UserRepository interface {
	Get(context context.Context, id int) core.User
}

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{
		repo,
	}
}

func (s *UserService) GetUserById(context context.Context, id int) core.User {
	return s.repo.Get(context, id)
}
