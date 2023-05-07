package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meziaris/devstore/internal/app/model"
	log "github.com/sirupsen/logrus"
)

type CategoryRepository struct {
	DB *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{DB: db}
}

func (cr *CategoryRepository) Create(category model.Category) error {
	var sqlStatement = `
		INSERT INTO categories (name, description)
		VALUES ($1, $2)
		`
	_, err := cr.DB.Exec(sqlStatement, category.Name, category.Description)
	if err != nil {
		return err
	}

	return nil
}

// get list category
func (cr *CategoryRepository) Browse(search model.BrowseCategory) ([]model.Category, error) {
	var (
		limit        = search.Limit
		offset       = limit * (search.Page - 1)
		categories   []model.Category
		sqlStatement = `
			SELECT id, name, description
			FROM categories
			LIMIT $1
			OFFSET $2
		`
	)

	rows, err := cr.DB.Queryx(sqlStatement, limit, offset)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - Browse : %w", err))
		return categories, err
	}

	for rows.Next() {
		var category model.Category
		err := rows.StructScan(&category)
		if err != nil {
			log.Error(fmt.Errorf("error CategoryRepository - Browse : %w", err))
		}
		categories = append(categories, category)
	}

	return categories, nil
}

// get detail category
func (cr *CategoryRepository) GetByID(id string) (model.Category, error) {
	var (
		sqlStatement = `
			SELECT id, name, description
			FROM categories
			WHERE id = $1
		`
		category model.Category
	)
	err := cr.DB.QueryRowx(sqlStatement, id).StructScan(&category)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - GetByID : %w", err))
		return category, err
	}

	return category, nil
}

// update article by id
func (cr *CategoryRepository) Update(category model.Category) error {
	var (
		sqlStatement = `
			UPDATE categories
			SET updated_at = NOW(),
				name = $2,
				description = $3
			WHERE id = $1
		`
	)

	result, err := cr.DB.Exec(sqlStatement, category.ID, category.Name, category.Description)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - UpdateByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

// delete article by id
func (cr *CategoryRepository) DeleteByID(id string) error {
	var (
		sqlStatement = `
			DELETE FROM categories
			WHERE id = $1
		`
	)

	result, err := cr.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error CategoryRepository - DeleteByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}
