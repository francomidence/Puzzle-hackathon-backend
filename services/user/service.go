package user

import (
	"context"
	"puzzle-hackathon-backend/models"
	"puzzle-hackathon-backend/repositories"
)

type UserService struct {
	UserRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{
		UserRepository: userRepository,
	}
}

func (s *UserService) CreateUser(ctx context.Context, payload models.User) (*models.User, error) {
	user, err := s.UserRepository.CreateUser(ctx, &payload)
	if err != nil {
		panic(err)
	}
	return user, nil
}

func (s *UserService) GetUsers(ctx context.Context) (*[]models.User, error) {
	users, err := s.UserRepository.GetUsers(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUser(ctx context.Context, payload models.User) (*models.User, error) {
	user, err := s.UserRepository.GetUser(ctx, &payload)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) UpdateUser(ctx context.Context, payload models.User) (*models.User, error) {
	patient, err := s.UserRepository.UpdateUser(ctx, &payload)
	if err != nil {
		panic(err)
	}
	return patient, nil
}

func (s *UserService) DeleteUser(ctx context.Context, payload models.User) (*models.User, error) {
	patient, err := s.UserRepository.DeleteUser(ctx, &payload)
	if err != nil {
		panic(err)
	}
	return patient, nil
}
