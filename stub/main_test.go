package main

import (
	"testing"
)

// ----------------------------------------------------------------------------
// A "stub" is an object that provides predefined answers to method calls.
// ----------------------------------------------------------------------------

// TestGetCityCoordinates tests the GetCityCoordinates function with both the real geo service and the stub geo service.

func TestGetCityCoordinates(t *testing.T) {
	// with real geo service:
	// It creates an instance of RealGeoService and passes it to GetCityCoordinates.
	// Since the real geo service returns dummy coordinates, we can only assert that the returned coordinates are within an acceptable range.
	t.Run("with real geo service", func(t *testing.T) {
		realGeoService := &RealGeoService{}
		lat, lng := GetCityCoordinates(realGeoService, "New York")
		// assert that the returned coordinates are within an acceptable range
		if lat < 37 || lat > 38 || lng < -123 || lng > -122 {
			t.Errorf("unexpected coordinates for New York: (%.4f, %.4f)", lat, lng)
		}
	})

	// with stub geo service:
	// It creates an instance of StubGeoService and passes it to GetCityCoordinates.
	// We assert that the returned coordinates match the predefined values for the specified city (New York in this case).
	t.Run("with stub geo service", func(t *testing.T) {
		stubGeoService := &StubGeoService{}
		lat, lng := GetCityCoordinates(stubGeoService, "New York")
		// assert that the returned coordinates match the predefined values
		expectedLat, expectedLng := 40.7128, -74.0060
		if lat != expectedLat || lng != expectedLng {
			t.Errorf("expected coordinates for New York: (%.4f, %.4f), but got: (%.4f, %.4f)",
				expectedLat, expectedLng, lat, lng)
		}
	})
}
