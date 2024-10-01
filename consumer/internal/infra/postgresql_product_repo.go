package infra

import (
	"database/sql"

	"github.com/guilherme-or/vivo-synq/consumer/internal/database"
	"github.com/guilherme-or/vivo-synq/consumer/internal/entity"
	"github.com/guilherme-or/vivo-synq/consumer/internal/repository"
)

const queryInsertProduct = "(id, status, product_name, subscription_type, start_date, end_date) VALUES ($1, $2, $3, $4, $5, $6);"
const queryUpdateProduct = "status = $1, product_name = $2, subscription_type = $3, start_date = $4, end_date = $5 WHERE id = $6"

type PostgreSQLProductRepository struct {
	db *sql.DB
}

func NewPostgreSQLProductRepository(conn *database.PostgreSQLConn) repository.ProductRepository {
	return &PostgreSQLProductRepository{db: conn.GetDatabase()}
}

func (r *PostgreSQLProductRepository) Insert(p *entity.Product) error {
	prefix := "INSERT INTO " + p.ProductType + "_products "

	res, err := r.db.Exec(
		prefix+queryInsertProduct,
		p.ID,
		p.Status,
		p.ProductName,
		p.SubscriptionType,
		p.StartDate,
		p.EndDate,
	)

	if err != nil {
		return err
	}

	if ra, err := res.RowsAffected(); err != nil || ra <= 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *PostgreSQLProductRepository) Update(before, after *entity.Product) error {
	prefix := "UPDATE " + before.ProductType + "_products SET "

	res, err := r.db.Exec(
		prefix+queryUpdateProduct,
		after.Status,
		after.ProductName,
		after.SubscriptionType,
		after.StartDate,
		after.EndDate,
		after.ID,
	)

	if err != nil {
		return err
	}

	if ra, err := res.RowsAffected(); err != nil || ra <= 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *PostgreSQLProductRepository) Delete(before *entity.Product) error {
	query := "DELETE FROM " + before.ProductType + "_products WHERE id = $1"

	res, err := r.db.Exec(
		query,
		before.ID,
	)

	if err != nil {
		return err
	}

	if ra, err := res.RowsAffected(); err != nil || ra <= 0 {
		return sql.ErrNoRows
	}

	return nil
}