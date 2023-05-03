package service

import "github.com/meziaris/devstore/internal/app/model"

type UserRepository interface {
	Create(user model.User) error
	GetByEmailAndUsername(email string, username string) (model.User, error)
	GetByEmail(email string) (model.User, error)
	GetByID(userID int) (model.User, error)
}
