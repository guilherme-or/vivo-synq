package repository

import (
	"database/sql"

	"github.com/guilherme-or/vivo-synq/kafka/internal/database"
	"github.com/guilherme-or/vivo-synq/kafka/internal/entity"
)

const queryInsertProduct = "(id, status, product_name, subscription_type, start_date, end_date) VALUES ($1, $2, $3, $4, $5, $6);"

type PostgreSQLProductRepository struct {
	db *sql.DB
}

func NewPostgreSQLProductRepository(conn *database.PostgreSQLConn) ProductRepository {
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

func (r *PostgreSQLProductRepository) Update(id int, p *entity.Product) error {
	return nil
}

func (r *PostgreSQLProductRepository) Delete(id int) error {
	return nil
}