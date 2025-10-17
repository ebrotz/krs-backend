package main

import (
	"net/http"

	"github.com/ebrotz/krs-backend/api"
	"github.com/ebrotz/krs-backend/internal/address"
)

func main() {
	// Use the address package to construct a ServerInterface-backed handler.
	m := api.Handler(address.NewAddressHandler())
	_ = http.ListenAndServe(":8080", m)
}
