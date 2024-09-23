package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/joshua468/team-budget-management/internal"
	"github.com/joshua468/team-budget-management/models"
)

func CreateTeamHandler(w http.ResponseWriter, r *http.Request) {
	var team models.Team
	if err := json.NewDecoder(r.Body).Decode(&team); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := internal.AddTeam(team); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(team)
}

func GetTeamsHandler(w http.ResponseWriter, r *http.Request) {
	teams, err := internal.GetTeams()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(teams)
}
