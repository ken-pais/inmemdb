package main

import (
	"testing"
	"time"
)

func TestDB(t *testing.T) {
	db := NewDB("FIFO")

	// Test: Add and Get an entry
	t.Run("Add and Get entry", func(t *testing.T) {
		db.AddEntry("key1", "value1", time.Hour)
		value, exists := db.Get("key1")
		if !exists || value != "value1" {
			t.Errorf("Expected 'value1', got '%s'", value)
		}
	})

	// Test: Retrieve a non-existent key
	t.Run("Get non-existent key", func(t *testing.T) {
		_, exists := db.Get("key_non_existent")
		if exists {
			t.Error("Expected key to be non-existent")
		}
	})

	

	// Test: Eviction after exceeding capacity (FIFO policy)
	t.Run("Eviction after exceeding capacity", func(t *testing.T) {
		// Set small capacity for testing
		db.Clear()
		db.Capacity = 2
		db.AddEntry("key1", "value1", time.Hour)
		db.AddEntry("key2", "value2", time.Hour)
		db.AddEntry("key3", "value3", time.Hour)

		// Verify eviction has occurred (FIFO)
		_, exists := db.Get("key1") // key1 should be evicted
		if exists {
			t.Error("Expected key1 to be evicted, but it was found")
		}

		// Verify remaining keys
		value, exists := db.Get("key2")
		if !exists || value != "value2" {
			t.Errorf("Expected 'value2' for key2, got '%s'", value)
		}

		value, exists = db.Get("key3")
		if !exists || value != "value3" {
			t.Errorf("Expected 'value3' for key3, got '%s'", value)
		}
	})
}
