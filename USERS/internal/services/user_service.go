package services

import (
	"USERS/internal/models"
	"USERS/internal/repository"
	"context"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetUserProfile(ctx context.Context, id int64) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *UserService) UpdateUserProfile(ctx context.Context, user *models.User) error {
	return s.repo.Update(ctx, user)
}
