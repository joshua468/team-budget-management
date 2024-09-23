package internal

import (
	"fmt"

	"github.com/joshua468/team-budget-management/config"
	"github.com/joshua468/team-budget-management/models"
)

// AddTeam creates a new team in the database.
func AddTeam(team models.Team) error {
	// Fix: Use the correct method "Create"
	if result := config.DB.Create(&team); result.Error != nil {
		return fmt.Errorf("failed to create team: %v", result.Error)
	}
	return nil
}

// GetTeams retrieves all teams along with their expenses.
func GetTeams() ([]models.Team, error) {
	var teams []models.Team

	// Fix: Pass the correct string to Preload (it expects a relation name, not a variable reference)
	if result := config.DB.Preload("Expenses").Find(&teams); result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve teams: %v", result.Error)
	}

	// Fix: Ensure that the condition in the if statement is a valid boolean condition
	if len(teams) == 0 {
		return nil, fmt.Errorf("no teams found")
	}

	return teams, nil
}
