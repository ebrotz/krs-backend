package places

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ebrotz/krs-backend/api"
	"github.com/ebrotz/krs-backend/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

// placeService represents the service layer that the application should use
// to interact with 'Place' resources.
type placeService struct {
	r repository
}

// make sure *placeService complies with api.StrictServerInterface
var _ api.StrictServerInterface = (*placeService)(nil)

func ptr[T any](val T) *T {
	return &val
}

// ListPlaces implements api.StrictServerInterface
func (s *placeService) ListPlaces(ctx context.Context, request api.ListPlacesRequestObject) (api.ListPlacesResponseObject, error) {
	log.Default().Println("ListPlaces")
	places, err := s.r.ListPlaces(ctx)

	if err != nil {
		return nil, err
	}

	return api.ListPlaces200JSONResponse(places), nil
}

// CreatePlace implements api.StrictServerInterface
func (s *placeService) CreatePlace(ctx context.Context, request api.CreatePlaceRequestObject) (api.CreatePlaceResponseObject, error) {
	panic("not implemented")
}

// DeletePlace implements api.StrictServerInterface
func (s *placeService) DeletePlace(ctx context.Context, request api.DeletePlaceRequestObject) (api.DeletePlaceResponseObject, error) {
	panic("not implemented")
}

// GetPlace implements api.StrictServerInterface
func (s *placeService) GetPlace(ctx context.Context, request api.GetPlaceRequestObject) (api.GetPlaceResponseObject, error) {
	panic("not implemented")
}

func (s *placeService) PatchPlace(ctx context.Context, request api.PatchPlaceRequestObject) (api.PatchPlaceResponseObject, error) {
	panic("not implemented")
}

// allowAllOrigins is a middleware that sets the 'Access-Control-Allow-Origins'
// header to '*'.
func allowAllOrigins(f api.StrictHandlerFunc, operationId string) api.StrictHandlerFunc {
	return func(ctx context.Context, w http.ResponseWriter, r *http.Request, request interface{}) (response interface{}, err error) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		return f(ctx, w, r, request)
	}
}

// NewPlaceHandler returns an http handler implementing the generated ServerInterface
// backed by placeService. The service methods are intentionally unimplemented stubs.
func NewPlaceHandler(ctx context.Context, config *config.Config) (http.Handler, error) {
	var r repository

	if config.DatabaseUrl != "" {
		pool, err := pgxpool.New(ctx, config.DatabaseUrl)

		if err != nil {
			return nil, fmt.Errorf("failed to create pgx pool: %w", err)
		}

		r = &postgresRepository{
			pool: pool,
		}
	} else {
		log.Default().Println("Using in-memory repository for Places")

		r = &inMemoryRepository{
			places: map[int]api.Place{
				0: {Name: ptr("Place 1"), Description: ptr("Seafood")},
				1: {Name: ptr("Place 2"), Description: ptr("Gastropub")},
				2: {Name: ptr("Place 3"), Description: ptr("Asian")},
			},
		}
	}

	service := &placeService{r: r}
	serverInterface := api.NewStrictHandler(service, []api.StrictMiddlewareFunc{allowAllOrigins})
	return api.Handler(serverInterface), nil
}
