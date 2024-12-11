package main

import "fmt"

// LRU eviction policy - Least Recently Used
type LRU struct{}

func (lru *LRU) Evict(db *DB) {
	fmt.Println("Implementing Least Recenty Used Policy")
}