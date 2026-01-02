package main

import (
	"net/http"

	"github.com/ebrotz/krs-backend/api"
	"github.com/ebrotz/krs-backend/internal/places"
)

func main() {
	m := api.Handler(places.NewPlaceHandler())
	_ = http.ListenAndServe(":8080", m)
}
