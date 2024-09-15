package repository

import "github.com/guilherme-or/vivo-synq/api/internal/entity"

type ProductRepository interface {
	Find(userID string) ([]entity.Product, error)              // Busca os produtos de um usuário por seu ID e retorna o objeto JSON em bytes
	FindInCache(userID string) ([]entity.Product, error)       // Busca os produtos de um usuário por seu ID no cache e retorna o objeto JSON em bytes
	SaveInCache(userID string, product []entity.Product) error // Salva os produtos de um usuário no cache
}
