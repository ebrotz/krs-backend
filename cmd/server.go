package main

import (
	"net/http"

	"github.com/ebrotz/krs-backend/internal/places"
)

func main() {
	m := http.NewServeMux()
	m.Handle("/", places.NewPlaceHandler())
	_ = http.ListenAndServe(":8080", m)
}
