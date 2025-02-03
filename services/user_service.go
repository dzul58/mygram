package services

import (
	"errors"
	"mygram/dto"
	"mygram/helpers"
	"mygram/models"
	"mygram/repositories"
)

type UserService interface {
	Register(input dto.UserRegisterRequest) (*dto.UserResponse, error)
	Login(input dto.UserLoginRequest) (string, error)
	Update(userID uint, input dto.UserRegisterRequest) (*dto.UserResponse, error)
	Delete(userID uint) error
}

type userService struct {
	userRepository repositories.UserRepository
}

func NewUserService(userRepository repositories.UserRepository) UserService {
	return &userService{userRepository}
}

func (s *userService) Register(input dto.UserRegisterRequest) (*dto.UserResponse, error) {
	hashedPassword, err := helpers.HashPassword(input.Password)
	if err != nil {
		return nil, err
	}

	user := &models.User{
		Username: input.Username,
		Email:    input.Email,
		Password: hashedPassword,
		Age:      input.Age,
	}

	if err := s.userRepository.Create(user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}, nil
}

func (s *userService) Login(input dto.UserLoginRequest) (string, error) {
	user, err := s.userRepository.FindByEmail(input.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	if !helpers.CheckPasswordHash(input.Password, user.Password) {
		return "", errors.New("invalid email or password")
	}

	token, err := helpers.GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (s *userService) Update(userID uint, input dto.UserRegisterRequest) (*dto.UserResponse, error) {
	user, err := s.userRepository.FindByID(userID)
	if err != nil {
		return nil, err
	}

	user.Username = input.Username
	user.Email = input.Email

	if err := s.userRepository.Update(user); err != nil {
		return nil, err
	}

	return &dto.UserResponse{
		ID:       user.ID,
		Username: user.Username,
		Email:    user.Email,
		Age:      user.Age,
	}, nil
}

func (s *userService) Delete(userID uint) error {
	return s.userRepository.Delete(userID)
}
