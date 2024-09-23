package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/joshua468/team-budget-management/internal"
	"github.com/joshua468/team-budget-management/models"

	"github.com/gorilla/mux"
)

func AddExpenseHandler(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	if err := json.NewDecoder(r.Body).Decode(&expense); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}
	if err := internal.AddExpense(expense); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(expense)
}

func GetExpensesByTeamHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	teamID, _ := strconv.Atoi(params["team_id"])
	expenses, err := internal.GetExpensesByTeam(uint(teamID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(expenses)
}
