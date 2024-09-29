package repository

import "github.com/guilherme-or/vivo-synq/consumer/internal/entity"

type TagRepository interface {
	Update(t *entity.Tags) error
}
