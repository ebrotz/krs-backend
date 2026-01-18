package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/ebrotz/krs-backend/internal/config"
	"github.com/ebrotz/krs-backend/internal/places"
)

func main() {
	c := config.ConfigFromEnv()
	ctx := context.Background()

	placeHandler, err := places.NewPlaceHandler(ctx, c)

	if err != nil {
		e := fmt.Errorf("failed to create place handler: %w", err)
		log.Fatal(e)
	}

	m := http.NewServeMux()
	m.Handle("/", placeHandler)

	fmt.Println("Starting krs-backend on port 8080")
	_ = http.ListenAndServe(":8080", m)
}
