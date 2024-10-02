package repository

import "github.com/guilherme-or/vivo-synq/consumer/internal/entity"

type ProductRepository interface {
	Insert(afterDTO *entity.ProductDTO) error
	Update(beforeDTO, afterDTO *entity.ProductDTO) error
	Delete(beforeDTO *entity.ProductDTO) error
}
