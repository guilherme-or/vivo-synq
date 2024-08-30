package repository

import "github.com/guilherme-or/vivo-synq/consumer/internal/entity"

type ProductRepository interface {
	Insert(p *entity.Product) error
	Update(id int, p *entity.Product) error
	Delete(id int, productType string) error
}
