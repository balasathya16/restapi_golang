package main

import (
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

// var teams []Team

// // Get all teams
// func getTeams(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(teams)
// }

// Get single team
func getFoods(w http.ResponseWriter, r *http.Request) {
}

func getFood(w http.ResponseWriter, r *http.Request) {
}

func createFood(w http.ResponseWriter, r *http.Request) {
}

func updateFood(w http.ResponseWriter, r *http.Request) {
}

func deleteFood(w http.ResponseWriter, r *http.Request) {
}

// Main function
func main() {
	// Init router
	router := mux.NewRouter()

	// // Hardcoded data - @todo: add database
	// teams = append(teams, team{ID: "1", Country: "India", League: "IPL"})
	// teams = append(teams, team{ID: "2", Country: "Canada", League: "GT20"})

	// Route handles & endpoints
	router.HandleFunc("/api/food", getFoods).Methods("GET")
	router.HandleFunc("/api/food/{id}", getFood).Methods("GET")
	router.HandleFunc("/api/food", createFood).Methods("POST")
	router.HandleFunc("/api/food/{id}", updateFood).Methods("PUT")
	router.HandleFunc("/api/food/{id}", deleteFood).Methods("DELETE")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", router))
}
