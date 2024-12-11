package main

import (
	"fmt"
	"time"
	"sync"
)

/*
In Memory Database/Cache
Features:
	- CRUD Operations 
		SET
		GET
		DEL
	Eviction Policies: Implements a FIFO eviction policy. You can extend it to add more policies like LRU, LFU, etc.
	TTL Support: Cache entries can expire based on a TTL (Time-to-Live).
	Clear Cache: Clears all data from the cache, resetting both the Data map and the Queue.
	Exposed APIs for testing purposes.
Design Patterns:
	Singleton Pattern: Ensures a single instance of the cache throughout the application lifecycle.
	Factory Pattern: Used for creating eviction policy struct by policy type
	Strategy Pattern: Used for implementing eviction policy struct by policy type
Data structures:
	Queue:
		- To keep track of the order of items inserted into DB
		+-----------+     +-----------+     +-----------+
		|    key1   | --> |    key2   | --> |    key3   |
		+-----------+     +-----------+     +-----------+
	Hashmap:
		+-----------+     +-----------+     +-----------+     +-----------+
		|   key1    | --> |   key2    | --> |   key3    | --> |   key4    |
		|   value1  |     |   value2  |     |   value3  |     |   value4  |
		+-----------+     +-----------+     +-----------+     +-----------+
*/

const capacity = 1024

type Entry struct {
	Value      string
	Expiration int64             // Unix timestamp for expiration time.Now().Unix()
}

// strategy pattern
type EvictionPolicy interface {
	Evict(db *DB)
}

type DB struct {
	Data     map[string]*Entry  // Cache data (key -> cache entry)
	Queue    []string           // Queue to maintain FIFO order
	Capacity int                // Maximum cache capacity
	Policy   EvictionPolicy     // Interface for implementing eviction policy
}

var instance *DB
var once sync.Once

// Factory pattern for eviction policies
type EvictionPolicyFactory interface {
	CreatePolicy() EvictionPolicy
}

type FIFOFactory struct{}

func (f *FIFOFactory) CreatePolicy() EvictionPolicy {
	return &FIFO{}
}

type LRUFactory struct{}

func (f *LRUFactory) CreatePolicy() EvictionPolicy {
	return &LRU{}
}

func GetEvictionPolicyFactory(policyType string) EvictionPolicyFactory {
	switch policyType {
	case "FIFO":
		return &FIFOFactory{}
	case "LRU":
		return &LRUFactory{}
	default:
		panic("Unknown eviction policy")
	}
}

// Singleton pattern - NewDB returns a singleton instance of the DB with a given capacity and eviction policy.
func NewDB(policy string) *DB {
	once.Do(func() {
		instance = &DB{
			Data:     make(map[string]*Entry),
			Queue:    []string{},
			Capacity: capacity,
			Policy:   GetEvictionPolicyFactory(policy).CreatePolicy(),
		}
	})
	return instance
}

// Get retrieves a value from the cache by key.
func (db *DB) Get(key string) (string, bool) {
	fmt.Println("Fetching entry:", "Key:", key)
	entry, exists := db.Data[key]
	if !exists {
		fmt.Println("key does not exist")
		return "", false
	}
	if entry.Expiration < time.Now().Unix() {
		fmt.Println("key expired")
		return "", false
	}
	fmt.Println("Fetched entry:", "Key:", key, "Value:", entry.Value)
	return entry.Value, true
}

// AddEntry adds or updates a key-value pair in the database.
func (db *DB) AddEntry(key, value string, ttl time.Duration) {
	// If the entry already exists, just update its value and expiration
	if entry, exists := db.Data[key]; exists {
		entry.Value = value
		entry.Expiration = time.Now().Add(ttl).Unix()
		fmt.Println("Updated entry:", "Key:", key, "Value:", value)
		return
	}

	// Otherwise, create a new entry
	newEntry := &Entry{
		Value:      value,
		Expiration: time.Now().Add(ttl).Unix(),
	}

	// Add to the map and the FIFO queue
	db.Data[key] = newEntry
	db.Queue = append(db.Queue, key)

	// If the cache exceeds the limit, evict
	if len(db.Data) > db.Capacity {
		fmt.Println("cache exceeds the limit")
		db.Policy.Evict(db)
	}
	fmt.Println("Created entry:", "Key:", key, "Value:", value)
	for _, val := range db.Queue {
		fmt.Println("\t\tqueue:", val)
	}
}

// Delete removes a key-value pair from the database.
func (db *DB) Delete(key string) {
	// Check if the key exists in Data
	if _, exists := db.Data[key]; !exists {
		fmt.Println("key does not exist")
		return
	}

	// Delete from the Data map
	delete(db.Data, key)

	// Remove from FIFO queue
	for i, k := range db.Queue {
		if k == key {
			db.Queue = append(db.Queue[:i], db.Queue[i+1:]...)
			break
		}
	}
	fmt.Println("Deleted entry:", "Key:", key)
}

func (db *DB) Clear() {
	db.Data = make(map[string]*Entry)
	db.Queue = []string{}
}
