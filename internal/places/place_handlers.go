package places

import (
	"context"

	"github.com/ebrotz/krs-backend/api"
)

type placeService struct {
	// TODO: add dependencies (db, logger, etc.)
}

// make sure *placeService complies with api.StrictServerInterface
var _ api.StrictServerInterface = (*placeService)(nil)

// ListPlaces implements api.StrictServerInterface
func (s *placeService) ListPlaces(ctx context.Context, request api.ListPlacesRequestObject) (api.ListPlacesResponseObject, error) {
	panic("not implemented")
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
func NewPlaceHandler() api.ServerInterface {
	return api.NewStrictHandler(&placeService{}, nil)
}
