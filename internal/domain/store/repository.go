package store

import (
	"context"
	"github.com/KbtuGophers/warehouse/internal/domain/inventory"
)

type Repository interface {
	Select(ctx context.Context) (dest []Entity, err error)
	Create(ctx context.Context, data Entity) (id string, err error)
	Get(ctx context.Context, id string) (dest Entity, err error)
	Update(ctx context.Context, id string, data Entity) (err error)
	Delete(ctx context.Context, id string) (err error)
	SelectInventories(ctx context.Context, id string) ([]inventory.Entity, error)
}
