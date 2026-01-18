package places

import (
	"encoding/json"
	"log"
	"net/http/httptest"
	"testing"

	"github.com/ebrotz/krs-backend/api"
	"github.com/ebrotz/krs-backend/internal/config"
)

func TestListPlaces(t *testing.T) {
	var places []api.Place
	h, err := NewPlaceHandler(t.Context(), &config.Config{})

	if err != nil {
		log.Fatal(err)
	}

	server := httptest.NewServer(h)
	defer server.Close()

	client := server.Client()
	resp, err := client.Get(server.URL + "/places")

	if err != nil {
		log.Fatal(err)
	}

	if resp == nil {
		t.Fatal("response is nil")
	}

	defer resp.Body.Close()

	if err := json.NewDecoder(resp.Body).Decode(&places); err != nil {
		t.Fatal(err)
	}

	if len(places) != 3 {
		t.Fatal("unexpected number of places")
	}
}
