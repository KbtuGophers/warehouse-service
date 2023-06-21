package warehouse

import "github.com/KbtuGophers/warehouse/internal/domain/store"

type Service struct {
	warehouseRepository store.Repository
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

func WithWarehouseRepository(warehouseRepository store.Repository) Configuration {
	return func(s *Service) error {
		s.warehouseRepository = warehouseRepository
		return nil
	}
}
