package places

import (
	"context"

	"github.com/ebrotz/krs-backend/api"
	"github.com/jackc/pgx/v5"
)

// repository is an abstraction layer of any persistence layer of places.
// It was written with the intention of hiding the underlying database
// connection, but may be used to abstract other persistence mechanisms.
type repository interface {
	ListPlaces(ctx context.Context) ([]api.Place, error)
	// GetPlace retrieves a place by its unique identifier
	GetPlace(id int) (api.Place, error)
}

var _ repository = (*postgresRepository)(nil)

// postgresRepository is an implementation of repository with an underlying
// pgx driver.
type postgresRepository struct {
	// database connection
	pg *pgx.Conn
}

func (p *postgresRepository) ListPlaces(ctx context.Context) ([]api.Place, error) {
	return nil, nil
}

func (p *postgresRepository) GetPlace(id int) (api.Place, error) {
	return api.Place{}, nil
}
