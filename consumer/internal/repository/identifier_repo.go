package repository

import "github.com/guilherme-or/vivo-synq/consumer/internal/entity"

type IdentifierRepository interface {
	Insert(after *entity.Identifier) error
	Update(before, after *entity.Identifier) error
	Delete(before *entity.Identifier) error
}
