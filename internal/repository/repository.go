package repository

import (
	//"github.com/KbtuGophers/warehouse/internal/repository/postgres"
	"github.com/KbtuGophers/warehouse/internal/domain/store"
	"github.com/KbtuGophers/warehouse/pkg/database"
)

type Configuration func(r *Repository) error

type Repository struct {
	postgres  *database.Database
	Warehouse store.Repository
	Inventory store.Repository
}

func New(configs ...Configuration) (r *Repository, err error) {
	// Create the repository
	r = &Repository{}

	// Apply all Configurations passed in
	for _, cfg := range configs {
		// Pass the repository into the configuration function
		if err = cfg(r); err != nil {
			return
		}
	}

	return
}

func (r *Repository) Close() {
	if r.postgres != nil {
		r.postgres.Client.Close()
	}
}

func WithPostgresStore(schema, dataSourceName string) Configuration {
	return func(r *Repository) (err error) {
		r.postgres, err = database.NewDatabase(schema, dataSourceName)
		if err != nil {
			return
		}

		if err = r.postgres.Migrate(); err != nil && err.Error() != "no change" {
			return
		}
		err = nil

		//r.Account = postgres.NewAccountRepository(r.postgres.Client)
		//r.Otp = postgres.NewOtpRepository(r.postgres.Client)

		return
	}
}
