package repository

import "github.com/guilherme-or/vivo-synq/consumer/internal/entity"

type DescriptionRepository interface {
	Update(d *entity.Description) error
}
