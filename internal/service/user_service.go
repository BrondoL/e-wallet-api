package service

import (
	"net/mail"

	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/dto"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/model"
	r "git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/internal/repository"
	"git.garena.com/sea-labs-id/batch-02/aulia-nabil/assignment-05-golang-backend/pkg/custom_error"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(input *dto.UserRequestBody) (*model.User, error)
}

type userService struct {
	userRepository r.UserRepository
}

type USConfig struct {
	UserRepository r.UserRepository
}

func NewUserService(c *USConfig) UserService {
	return &userService{
		userRepository: c.UserRepository,
	}
}

func (s *userService) CreateUser(input *dto.UserRequestBody) (*model.User, error) {
	_, err := mail.ParseAddress(input.Email)
	if err != nil {
		return &model.User{}, &custom_error.NotValidEmailError{}
	}

	user, err := s.userRepository.FindByEmail(input.Email)
	if err != nil {
		return user, err
	}
	if user.ID != 0 {
		return user, &custom_error.UserAlreadyExistsError{}
	}

	user.Name = input.Name
	user.Email = input.Email
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.MinCost)
	if err != nil {
		return user, err
	}
	user.Password = string(passwordHash)

	newUser, err := s.userRepository.Save(user)
	if err != nil {
		return newUser, err
	}
	return newUser, nil
}
