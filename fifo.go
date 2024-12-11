// inmemdb/eviction/fifo.go
package main

import (
	"fmt"
)

// FIFO eviction policy - First In, First Out
type FIFO struct{}

// Evict removes the oldest entry (the first inserted).
func (fifo *FIFO) Evict(db *DB) {
	if len(db.Queue) < 1 {
		fmt.Println("Queue is empty, no eviction needed.")
		return
	}
	fmt.Println("Implementing FIFO eviction.")
	// Evict the oldest item (front of the queue)
	oldestKey := db.Queue[0]
	db.Delete(oldestKey)
	if len(db.Queue) > 0 {
		// Remove the evicted key from the queue (after evicting the oldest)
		db.Queue = db.Queue[1:]
	} else {
		// If the queue is empty, no further slicing is necessary
		db.Queue = []string{}
	}
}
