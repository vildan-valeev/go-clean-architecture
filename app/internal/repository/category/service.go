package user

import (
	"github.com/vildan-valeev/go-clean-architecture/internal/repository"
)

type Repository struct {
	db repository.DB
	rs repository.RedisCache
}

func New(db repository.DB) *Repository {
	return &Repository{db: db}
}
