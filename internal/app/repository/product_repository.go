package repository

import (
	"errors"
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/meziaris/devstore/internal/app/model"
	log "github.com/sirupsen/logrus"
)

type ProductRepository struct {
	DB *sqlx.DB
}

func NewProducRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{DB: db}
}

func (r *ProductRepository) Create(product model.Product) (productID int, err error) {
	var (
		id           int
		sqlStatement = `
			INSERT INTO products (name,description,currency,total_stock,is_active,category_id)
			VALUES (:name, :description, :currency, :total_stock, :is_active, :category_id)
			RETURNING id
		`
	)

	stmt, err := r.DB.PrepareNamed(sqlStatement)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - PrepareNamed : %w", err))
		return 0, err
	}

	err = stmt.Get(&id, product)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Create : %w", err))
		return 0, err
	}

	return id, nil
}

func (r *ProductRepository) Browse(search model.BrowseProduct) ([]model.Product, error) {
	var (
		limit        = search.Limit
		offset       = limit * (search.Page - 1)
		products     []model.Product
		sqlStatement = `
			SELECT id, name, description, currency, total_stock, is_active, category_id, image_url
			FROM products
			LIMIT $1
			OFFSET $2
		`
	)

	rows, err := r.DB.Queryx(sqlStatement, limit, offset)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Browse : %w", err))
		return products, err
	}

	for rows.Next() {
		product := model.Product{}
		if err := rows.StructScan(&product); err != nil {
			log.Error(fmt.Errorf("error ProductRepository - Browse : %w", err))
		}
		products = append(products, product)
	}

	return products, nil
}

func (r *ProductRepository) GetByID(id string) (model.Product, error) {
	var (
		product      model.Product
		sqlStatement = `
			SELECT id, name, description, currency, total_stock, is_active, category_id, image_url
			FROM products
			WHERE id = $1
		`
	)

	if err := r.DB.QueryRowx(sqlStatement, id).StructScan(&product); err != nil {
		log.Error(fmt.Errorf("error ProductRepository - GetByID : %w", err))
		return product, err
	}

	return product, nil
}

func (r *ProductRepository) Update(product model.Product) error {
	var (
		sqlStatement = `
			UPDATE products
			SET updated_at = NOW(),
				name = $2,
				description = $3,
				currency = $4,
				total_stock = $5,
				is_active = $6,
				category_id = $7
			WHERE id = $1
		`
	)

	result, err := r.DB.Exec(sqlStatement,
		product.ID,
		product.Name,
		product.Description,
		product.Currency,
		product.TotalStock,
		product.IsActive,
		product.CategoryID,
	)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - Update : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

func (r *ProductRepository) UpdateImageURL(id int, imageURL string) error {
	var (
		sqlStatement = `
			UPDATE products
			SET updated_at = NOW(), image_url = $2
			WHERE id = $1
		`
	)

	result, err := r.DB.Exec(sqlStatement, id, imageURL)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - UpdateImageURL : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}

func (r *ProductRepository) DeleteByID(id string) error {
	var (
		sqlStatement = `
			DELETE FROM products
			WHERE id = $1
		`
	)

	result, err := r.DB.Exec(sqlStatement, id)
	if err != nil {
		log.Error(fmt.Errorf("error ProductRepository - DeleteByID : %w", err))
		return err
	}

	totalAffected, _ := result.RowsAffected()
	if totalAffected <= 0 {
		return errors.New("no record affected")
	}

	return nil
}
