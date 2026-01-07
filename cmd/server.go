package main

import (
	"fmt"
	"net/http"

	"github.com/ebrotz/krs-backend/internal/places"
)

func main() {
	m := http.NewServeMux()
	m.Handle("/", places.NewPlaceHandler())
	fmt.Println("Starting krs-backend on port 8080")
	_ = http.ListenAndServe(":8080", m)
}
