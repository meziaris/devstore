package service

import (
	"errors"

	"github.com/meziaris/devstore/internal/app/model"
	"github.com/meziaris/devstore/internal/app/schema"
	"github.com/meziaris/devstore/internal/pkg/reason"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	Create(user model.User) error
	GetByEmailAndUsername(email string, username string) (model.User, error)
}

type RegistrationService struct {
	userRepo UserRepository
}

func NewRegistrationService(userRepo UserRepository) *RegistrationService {
	return &RegistrationService{
		userRepo: userRepo,
	}
}

func (s *RegistrationService) Register(req *schema.RegisterReq) error {
	existingUser, _ := s.userRepo.GetByEmailAndUsername(req.Email, req.Username)
	if existingUser.ID > 0 {
		return errors.New(reason.UserAlreadyExist)
	}
	password, _ := s.hashPassword(req.Password)
	inserData := model.User{
		Username:       req.Username,
		Email:          req.Email,
		HashedPassword: password,
	}
	if err := s.userRepo.Create(inserData); err != nil {
		return errors.New(reason.RegisterFailed)
	}

	return nil
}

func (s *RegistrationService) hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(bytes), nil
}
