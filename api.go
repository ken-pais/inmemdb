package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

func StartServer(db *DB) {
	http.HandleFunc("/get/", handleGet(db))
	http.HandleFunc("/set/", handleSet(db))
	http.HandleFunc("/delete/", handleDelete(db))
	fmt.Println("Starting server on :8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

func handleGet(db *DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path[len("/get/"):]
		if key == "" {
			http.Error(w, "Key is missing", http.StatusBadRequest)
			return
		}
		value, exists := db.Get(key)
		if exists {
			w.Header().Set("Content-Type", "application/json")
			err := json.NewEncoder(w).Encode(map[string]string{"key": key, "value": value})
			if err != nil {
				http.Error(w, "Error encoding response", http.StatusInternalServerError)
				return
			}
		} else {
			http.Error(w, "Key not found", http.StatusNotFound)
		}
	}
}

func handleSet(db *DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path[len("/set/"):]
		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, "Error reading request body", http.StatusBadRequest)
			return
		}
		var data map[string]string
		err = json.Unmarshal(body, &data)
		if err != nil || data["value"] == "" {
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}
		value := data["value"]
		ttl := time.Hour
		db.AddEntry(key, value, ttl)
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Key added", "key": key})
	}
}

func handleDelete(db *DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		key := r.URL.Path[len("/delete/"):]
		db.Delete(key)
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{"message": "Key deleted", "key": key})
	}
}
