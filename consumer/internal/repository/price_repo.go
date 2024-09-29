package repository

import "github.com/guilherme-or/vivo-synq/consumer/internal/entity"

type PriceRepository interface {
	Update(p *entity.Price) error
}
