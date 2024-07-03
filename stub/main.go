package main

import "fmt"

// ----------------------------------------------------------------------------
// A "stub" is an object that provides predefined answers to method calls.
// ----------------------------------------------------------------------------

// We define the GeoService interface and two structs: RealGeoService and StubGeoService.

// The RealGeoService simulates a real geocoding service that would typically make an API call to retrieve the actual coordinates for a given address.
// For simplicity, it returns dummy coordinates in this example.

// The StubGeoService provides predefined coordinates for specific city names.

// The GetCityCoordinates function takes a GeoService interface and a city string.
// It calls the GetCoordinates method of the provided geo service to retrieve the coordinates for the given city.

// We demonstrate the usage of both the real geo service and the stub geo service.
// When using the RealGeoService, it returns the dummy coordinates.
// When using the StubGeoService, it returns the predefined coordinates for the specified city.

// The purpose of the stub geo service in this example is to provide predefined answers for specific inputs.
// Instead of relying on a real geocoding service, which might be slow or have external dependencies, we use a stub that returns hardcoded coordinates for known cities.
// This allows us to test the behaviour of GetCityCoordinates in a controlled and predictable manner.

// In real-world usage, you would typically use the RealGeoService (or any other concrete geocoding service implementation) to retrieve the actual coordinates.
// The stub geo service is primarily used in testing or in situations where you want to control the behaviour of the geocoding service for specific inputs.

type GeoService interface {
	GetCoordinates(address string) (float64, float64)
}

type RealGeoService struct{}

func (gs *RealGeoService) GetCoordinates(address string) (float64, float64) {
	// simulating a real geocoding service
	// this would typically make an API call to a geocoding service
	// and return the actual coordinates for the given address.
	// for simplicity, we'll return dummy coordinates here.
	return 37.7749, -122.4194
}

// this is the "stub" an object that provides predefined answers to method calls
type StubGeoService struct{}

func (gs *StubGeoService) GetCoordinates(address string) (float64, float64) {
	// return predefined coordinates for specific addresses
	if address == "New York" {
		return 40.7128, -74.0060
	} else if address == "London" {
		return 51.5074, -0.1278
	}
	return 0, 0
}

func GetCityCoordinates(geoService GeoService, city string) (float64, float64) {
	return geoService.GetCoordinates(city)
}

func main() {
	city := "New York"

	// using a real geo service
	realGeoService := &RealGeoService{}
	lat, lng := GetCityCoordinates(realGeoService, city)
	fmt.Printf("Real GeoService - Coordinates for %s: (%.4f, %.4f)\n", city, lat, lng)

	// using a stub geo service (! typically only used in the tests)
	stubGeoService := &StubGeoService{}
	lat, lng = GetCityCoordinates(stubGeoService, city)
	fmt.Printf("Stub GeoService - Coordinates for %s: (%.4f, %.4f)\n", city, lat, lng)
}
