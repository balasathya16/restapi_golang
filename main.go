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

// // Get single team
// func getTeam(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	params := mux.Vars(r) // Gets params
// 	// Loop through teams and find one team with the id from the params
// 	for _, item := range teams {
// 		if item.ID == params["id"] {
// 			json.NewEncoder(w).Encode(item)
// 			return
// 		}
// 	}
// 	json.NewEncoder(w).Encode(&Team{})
// }

// // Add new team
// func createTeam(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	var team Team
// 	_ = json.NewDecoder(r.Body).Decode(&team)
// 	team.ID = strconv.Itoa(rand.Intn(100000000)) // sample ID
// 	teams = append(teams, team)
// 	json.NewEncoder(w).Encode(team)
// }

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
