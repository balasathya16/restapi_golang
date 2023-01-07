package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Teams struct

type Team struct {
	ID      string `json:"id"`
	Country string `json:"country"`
	League  string `json:"league"`
}

var teams []Team

// Get all teams
func getTeams(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(teams)
}

// Get single team
func getTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Gets params
	// Loop through teams and find one team with the id from the params
	for _, item := range teams {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Team{})
}

// Add new team
func createTeam(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var team Team
	_ = json.NewDecoder(r.Body).Decode(&team)
	team.ID = strconv.Itoa(rand.Intn(100000000)) // sample ID
	teams = append(teams, team)
	json.NewEncoder(w).Encode(team)
}

// Main function
func main() {
	// Init router
	r := mux.NewRouter()

	// Hardcoded data - @todo: add database
	teams = append(teams, team{ID: "1", Country: "India", League: "IPL"})
	teams = append(teams, team{ID: "2", Country: "Canada", League: "GT20"})

	// Route handles & endpoints
	r.HandleFunc("/teams", getTeams()).Methods("GET")
	r.HandleFunc("/teams/{id}", getTeam()).Methods("GET")
	r.HandleFunc("/teams", createTeam()).Methods("POST")

	// Start server
	log.Fatal(http.ListenAndServe(":8000", r))
}

// Request sample
// {
// 	"isbn":"4545454",
// 	"title":"Book Three",
// 	"author":{"firstname":"Harry","lastname":"White"}
// }
