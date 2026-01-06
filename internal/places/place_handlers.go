package places

import (
	"context"
	"net/http"

	"github.com/ebrotz/krs-backend/api"
)

type placeService struct {
	// TODO: add dependencies (db, logger, etc.)
}

// make sure *placeService complies with api.StrictServerInterface
var _ api.StrictServerInterface = (*placeService)(nil)

func ptr[T any](val T) *T {
	return &val
}

// ListPlaces implements api.StrictServerInterface
func (s *placeService) ListPlaces(ctx context.Context, request api.ListPlacesRequestObject) (api.ListPlacesResponseObject, error) {
	// TODO This should be retrieved from the database.
	places := []api.Place{
		{Name: ptr("Place 1")},
		{Name: ptr("Place 2")},
		{Name: ptr("Place 3")},
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

// NewPlaceHandler returns an http handler implementing the generated ServerInterface
// backed by placeService. The service methods are intentionally unimplemented stubs.
func NewPlaceHandler() http.Handler {
	service := &placeService{}
	serverInterface := api.NewStrictHandler(service, nil)
	return api.Handler(serverInterface)
}
