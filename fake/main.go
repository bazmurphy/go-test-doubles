package main

import "fmt"

// A "fake" is a lightweight implementation of a real object that takes shortcuts to simplify its behaviour.

// We define the Database interface and two structs: RealDatabase and FakeDatabase.

// The RealDatabase simulates a real database implementation, where saving and retrieving data would typically involve interacting with an actual database. For simplicity, it just prints messages to indicate the operations.

// The FakeDatabase is a lightweight implementation of the Database interface that uses an in-memory map to store key-value pairs.
// It provides simplified implementations of the Save and Get methods.

// The SaveUserData function takes a Database interface and userID and userData strings.
// It calls the Save method of the provided database to store the user data.

// The GetUserData function takes a Database interface and a userID string.
// It calls the Get method of the provided database to retrieve the user data.

// We demonstrate the usage of both the real database and the fake database.
// When using the RealDatabase, it simulates saving and retrieving data from a real database.
// When using the FakeDatabase, it stores and retrieves data using the in-memory map.

// The purpose of the fake database in this example is to provide a lightweight implementation of the database that simplifies its behaviour.
// Instead of relying on a real database, which might require setup, teardown, and external dependencies, we use an in-memory map to simulate the database operations.
// This allows us to test the behaviour of SaveUserData and GetUserData without the need for an actual database.

// In real-world usage, you would typically use the RealDatabase (or any other concrete database implementation) to interact with an actual database.
// The fake database is primarily used in testing or in situations where you want to simplify the database behaviour for testing or development purposes.

type Database interface {
	Save(key, value string)
	Get(key string) (string, bool)
}

type RealDatabase struct{}

func (db *RealDatabase) Save(key, value string) {
	// simulating saving data to a real database
	fmt.Printf("saving data: key=%s, value=%s\n", key, value)
	// in a real implementation, this would store the data in an actual database
}

func (db *RealDatabase) Get(key string) (string, bool) {
	// simulating retrieving data from a real database
	fmt.Printf("retrieving data: key=%s\n", key)
	// in a real implementation, this would fetch the data from an actual database
	// but for simplicity, we just return a value and true to indicate success
	return "some user data", true
}

type FakeDatabase struct {
	data map[string]string
}

func NewFakeDatabase() *FakeDatabase {
	return &FakeDatabase{
		data: make(map[string]string),
	}
}

func (db *FakeDatabase) Save(key, value string) {
	db.data[key] = value
}

func (db *FakeDatabase) Get(key string) (string, bool) {
	value, exists := db.data[key]
	return value, exists
}

func SaveUserData(db Database, userID, userData string) {
	db.Save(userID, userData)
}

func GetUserData(db Database, userID string) (string, bool) {
	return db.Get(userID)
}

func main() {
	userID := "user123"
	userData := "some user data"

	// using a real database
	realDB := &RealDatabase{}
	SaveUserData(realDB, userID, userData)
	data, exists := GetUserData(realDB, userID)
	fmt.Printf("Real Database - User data for %s: %s (exists: %v)\n", userID, data, exists)

	// using a fake database (! typically only used in the tests)
	fakeDB := NewFakeDatabase()
	SaveUserData(fakeDB, userID, userData)
	data, exists = GetUserData(fakeDB, userID)
	fmt.Printf("Fake Database - User data for %s: %s (exists: %v)\n", userID, data, exists)
}
