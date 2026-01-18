package places

import (
	"context"
	"fmt"

	"github.com/ebrotz/krs-backend/api"
	v1 "github.com/ebrotz/krs-backend/internal/data/v1"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

// repository is an abstraction layer of any persistence layer of places.
// It was written with the intention of hiding the underlying database
// connection, but may be used to abstract other persistence mechanisms.
type repository interface {
	ListPlaces(ctx context.Context) ([]api.Place, error)
	// GetPlace retrieves a place by its unique identifier
	GetPlace(id int) (*api.Place, error)
}

var _ repository = (*postgresRepository)(nil)

// postgresRepository is an implementation of repository with an underlying
// pgx driver.
type postgresRepository struct {
	// database connection
	pool *pgxpool.Pool
}

func (p *postgresRepository) ListPlaces(ctx context.Context) ([]api.Place, error) {
	var ret []api.Place
	rows, err := p.pool.Query(ctx, "select * from places")

	if err != nil {
		return nil, fmt.Errorf("failed to query from db: %w", err)
	}

	places, err := pgx.CollectRows(rows, pgx.RowToStructByName[v1.Place])

	if err != nil {
		return nil, fmt.Errorf("pg repository: failed to collect rows %w", err)
	}

	for _, p := range places {
		ret = append(ret, api.Place{
			// TODO ID
			Description: &p.Description,
			Name:        &p.Name,
		})
	}

	return ret, nil
}

func (p *postgresRepository) GetPlace(id int) (*api.Place, error) {
	return nil, nil
}

// inMemoryRepository is a repository implementation that simply keeps a collection
// of places in a map.
type inMemoryRepository struct {
	places map[int]api.Place
}

var _ repository = (*inMemoryRepository)(nil)

func (r *inMemoryRepository) ListPlaces(ctx context.Context) ([]api.Place, error) {
	var ret []api.Place

	for _, p := range r.places {
		ret = append(ret, p)
	}
	return ret, nil
}

func (r *inMemoryRepository) GetPlace(id int) (*api.Place, error) {
	p, exists := r.places[id]

	if !exists {
		return nil, fmt.Errorf("in memory places repository: %d not in map", id)
	}

	return &p, nil
}
