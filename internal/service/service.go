package service

import (
	"github.com/KbtuGophers/warehouse/internal/service/inventory"
	"github.com/KbtuGophers/warehouse/internal/service/warehouse"
)

type Service struct {
	WarehouseService warehouse.Service
	InventoryService inventory.Service
}

func NewService(WarehouseService warehouse.Service, InventoryService inventory.Service) *Service {
	return &Service{
		WarehouseService: WarehouseService,
		InventoryService: InventoryService,
	}
}
