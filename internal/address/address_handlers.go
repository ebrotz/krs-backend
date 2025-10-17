package address

import (
	"context"

	"github.com/ebrotz/krs-backend/api"
)

type addressService struct {
	// TODO: add dependencies (db, logger, etc.)
}

// make sure *addressService complies with api.StrictServerInterface
var _ api.StrictServerInterface = (*addressService)(nil)

// ListAddresses implements api.StrictServerInterface
func (s *addressService) ListAddresses(ctx context.Context, request api.ListAddressesRequestObject) (api.ListAddressesResponseObject, error) {
	panic("not implemented")
}

// CreateAddress implements api.StrictServerInterface
func (s *addressService) CreateAddress(ctx context.Context, request api.CreateAddressRequestObject) (api.CreateAddressResponseObject, error) {
	panic("not implemented")
}

// DeleteAddress implements api.StrictServerInterface
func (s *addressService) DeleteAddress(ctx context.Context, request api.DeleteAddressRequestObject) (api.DeleteAddressResponseObject, error) {
	panic("not implemented")
}

// GetAddress implements api.StrictServerInterface
func (s *addressService) GetAddress(ctx context.Context, request api.GetAddressRequestObject) (api.GetAddressResponseObject, error) {
	panic("not implemented")
}

// UpdateAddress implements api.StrictServerInterface
func (s *addressService) UpdateAddress(ctx context.Context, request api.UpdateAddressRequestObject) (api.UpdateAddressResponseObject, error) {
	panic("not implemented")
}

// ListRestaurants implements api.StrictServerInterface
func (s *addressService) ListRestaurants(ctx context.Context, request api.ListRestaurantsRequestObject) (api.ListRestaurantsResponseObject, error) {
	panic("not implemented")
}

// CreateRestaurant implements api.StrictServerInterface
func (s *addressService) CreateRestaurant(ctx context.Context, request api.CreateRestaurantRequestObject) (api.CreateRestaurantResponseObject, error) {
	panic("not implemented")
}

// DeleteRestaurant implements api.StrictServerInterface
func (s *addressService) DeleteRestaurant(ctx context.Context, request api.DeleteRestaurantRequestObject) (api.DeleteRestaurantResponseObject, error) {
	panic("not implemented")
}

// GetRestaurant implements api.StrictServerInterface
func (s *addressService) GetRestaurant(ctx context.Context, request api.GetRestaurantRequestObject) (api.GetRestaurantResponseObject, error) {
	panic("not implemented")
}

// UpdateRestaurant implements api.StrictServerInterface
func (s *addressService) UpdateRestaurant(ctx context.Context, request api.UpdateRestaurantRequestObject) (api.UpdateRestaurantResponseObject, error) {
	panic("not implemented")
}

// NewAddressHandler returns an http handler implementing the generated ServerInterface
// backed by addressService. The service methods are intentionally unimplemented stubs.
func NewAddressHandler() api.ServerInterface {
	return api.NewStrictHandler(&addressService{}, nil)
}
