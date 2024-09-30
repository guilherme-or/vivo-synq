package repository

import "github.com/guilherme-or/vivo-synq/consumer/internal/entity"

type DescriptionRepository interface {
	Insert(after *entity.Description) error
	Update(before, after *entity.Description) error
	Delete(before *entity.Description) error
}
