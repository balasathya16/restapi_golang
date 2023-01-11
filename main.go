package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

// Food struct

type Food struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Cuisine string `json:"cuisine"`
}

var foods []Food

// Get all food
func getFoods(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foods)
}

// Get single food

func getFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // get parameters
	// loop through the foods and find using ID
	for _, item := range foods {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Food{})
}

// Create food

func createFood(w http.ResponseWriter, r *http.Request) {
}

// Update food

func updateFood(w http.ResponseWriter, r *http.Request) {
}

// Delete food

func deleteFood(w http.ResponseWriter, r *http.Request) {
}

// Main function
func main() {
	// Init router
	router := mux.NewRouter()

	// Hardcoded data - @todo: add database
	foods = append(foods, Food{ID: "1", Name: "Pasta", Cuisine: "Italian"})
	foods = append(foods, Food{ID: "2", Name: "Idly", Cuisine: "Indian"})

	router.HandleFunc("/api/food", getFoods).Methods("GET")
	router.HandleFunc("/api/food/{id}", getFood).Methods("GET")
	router.HandleFunc("/api/food", createFood).Methods("POST")
	router.HandleFunc("/api/food/{id}", updateFood).Methods("PUT")
	router.HandleFunc("/api/food/{id}", deleteFood).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
