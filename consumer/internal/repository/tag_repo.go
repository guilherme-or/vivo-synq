package repository

import "github.com/guilherme-or/vivo-synq/consumer/internal/entity"

type TagRepository interface {
	Insert(after *entity.Tag) error
	Update(before, after *entity.Tag) error
	Delete(before *entity.Tag) error
}
