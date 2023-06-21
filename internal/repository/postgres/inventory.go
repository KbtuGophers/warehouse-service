package postgres

import (
	"context"
	"github.com/KbtuGophers/warehouse/internal/domain/inventory"
	"github.com/jmoiron/sqlx"
)

type InventoryRepository struct {
	db *sqlx.DB
}

func (r *InventoryRepository) Select(ctx context.Context) (dest []inventory.Entity, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *InventoryRepository) Create(ctx context.Context, data inventory.Entity) (id string, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *InventoryRepository) Get(ctx context.Context, id string) (dest inventory.Entity, err error) {
	//TODO implement me
	panic("implement me")
}

func (r *InventoryRepository) Update(ctx context.Context, id string, data inventory.Entity) (err error) {
	//TODO implement me
	panic("implement me")
}
