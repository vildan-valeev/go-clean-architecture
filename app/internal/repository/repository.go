package repository

import "context"

type Repositories struct {
	Category *CategoryRepository
	Item     *ItemRepository
}

type Deps struct {
	PG PostgresDB
	MG MongoDB
	RS RedisDB

	ItemMongoCollection     string
	CategoryMongoCollection string
}

func NewRepositories(ctx context.Context, deps Deps) *Repositories {
	c := NewCategoryRepo(deps.PG, deps.MG.Collection(ctx, deps.CategoryMongoCollection))
	i := NewItemRepo(deps.PG, deps.RS, deps.MG.Collection(ctx, deps.ItemMongoCollection))

	return &Repositories{
		Category: c,
		Item:     i,
	}
}
