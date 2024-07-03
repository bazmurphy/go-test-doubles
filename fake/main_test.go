package main

import (
	"testing"
)

// ----------------------------------------------------------------------------
// A "fake" is a lightweight implementation of a real object that takes shortcuts to simplify its behaviour.
// ----------------------------------------------------------------------------

// TestSaveUserData tests the SaveUserData function with both the real database and the fake database.

func TestSaveUserData(t *testing.T) {
	// "with real database":
	// creates an instance of RealDatabase and passes it to SaveUserData.
	// no assertions are needed here because the real database is used, and its behaviour is not being tested.
	t.Run("with real database", func(t *testing.T) {
		realDB := &RealDatabase{}
		userID := "user123"
		userData := "some user data"
		SaveUserData(realDB, userID, userData)
	})

	// "with fake database":
	// creates an instance of FakeDatabase and passes it to SaveUserData.
	// after saving the user data, it retrieves the data using GetUserData and asserts that the data exists and matches the expected value.
	t.Run("with fake database", func(t *testing.T) {
		fakeDB := NewFakeDatabase()
		userID := "user123"
		userData := "some user data"
		SaveUserData(fakeDB, userID, userData)
		retrievedData, exists := GetUserData(fakeDB, userID)
		if !exists {
			t.Errorf("expected user data to exist for user %s", userID)
		}
		if retrievedData != userData {
			t.Errorf("expected user data: %s, but got: %s", userData, retrievedData)
		}
	})
}
