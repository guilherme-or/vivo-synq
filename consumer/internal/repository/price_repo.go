package repository

import "github.com/guilherme-or/vivo-synq/consumer/internal/entity"

type PriceRepository interface {
	Insert(after *entity.Price) error
	Update(before, after *entity.Price) error
	Delete(before *entity.Price) error
}
