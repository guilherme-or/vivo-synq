package repository

import "github.com/guilherme-or/vivo-synq/consumer/internal/entity"

type IdentifierRepository interface {
	Update(i *entity.Identifiers) error
}
