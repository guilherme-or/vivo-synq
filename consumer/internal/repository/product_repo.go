package repository

import "github.com/guilherme-or/vivo-synq/consumer/internal/entity"

type ProductRepository interface {
	Insert(after *entity.Product) error
	Update(before, after *entity.Product) error
	Delete(before *entity.Product) error
}
