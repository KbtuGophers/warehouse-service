package handler

import (
	"github.com/KbtuGophers/warehouse/internal/config"
	"github.com/KbtuGophers/warehouse/internal/service/inventory"
	"github.com/KbtuGophers/warehouse/internal/service/warehouse"
	"github.com/KbtuGophers/warehouse/pkg/server/router"
	"github.com/go-chi/chi/v5"
)

type Configuration func(h *Handler) error

type Dependencies struct {
	Configs   config.Config
	Warehouse *warehouse.Service
	Inventory *inventory.Service
}

type Handler struct {
	//Service *service.Service

	dependencies Dependencies
	HTTP         *chi.Mux
}

func New(d Dependencies, configs ...Configuration) (h *Handler, err error) {
	// Create the handler
	h = &Handler{
		dependencies: d,
	}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the service into the configuration function
		if err = cfg(h); err != nil {
			return
		}
	}

	return
}

func WithHTTPHandler() Configuration {
	return func(h *Handler) (err error) {
		h.HTTP = router.New()

		//docs.SwaggerInfo.BasePath = "/api/v1"
		//docs.SwaggerInfo.Host = h.dependencies.Configs.HTTP.Host
		//docs.SwaggerInfo.Schemes = []string{h.dependencies.Configs.HTTP.Schema}
		//docs.SwaggerInfo.Title = "Account Service"
		return
	}
}
