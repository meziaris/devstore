package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meziaris/devstore/internal/app/model"
	log "github.com/sirupsen/logrus"
)

type UserRepository struct {
	DB *sqlx.DB
}

func NewUserRepository(db *sqlx.DB) *UserRepository {
	return &UserRepository{DB: db}
}

// create user
func (r *UserRepository) Create(user model.User) error {
	var (
		sqlStatement = `
			INSERT INTO users (username, email, hashed_password)
			VALUES ($1, $2, $3)
		`
	)

	_, err := r.DB.Exec(sqlStatement, user.Username, user.Email, user.HashedPassword)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - Create : %w", err))
		return err
	}

	return nil
}

// get detail user by email and username
func (r *UserRepository) GetByEmailAndUsername(email string, username string) (model.User, error) {
	var (
		sqlStatement = `
			SELECT id, username, email
			FROM users
			WHERE email = $1 AND username = $2
			LIMIT 1
		`
		user model.User
	)

	err := r.DB.QueryRowx(sqlStatement, email, username).StructScan(&user)
	if err != nil {
		log.Error(fmt.Errorf("error UserRepository - GetByEmailAndUsername : %w", err))
		return user, err
	}

	return user, nil
}
