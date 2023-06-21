package inventory

import (
	"github.com/KbtuGophers/warehouse/internal/domain/inventory"
)

type Service struct {
	inventoryRepository inventory.Repository
}

type Configuration func(s *Service) error

func New(configs ...Configuration) (s *Service, err error) {
	// Create the service
	s = &Service{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		if err = cfg(s); err != nil {
			return
		}
	}
	return
}

func WithInventoryRepository(inventoryRepository inventory.Repository) Configuration {
	return func(s *Service) error {
		s.inventoryRepository = inventoryRepository
		return nil
	}
}
