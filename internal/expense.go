package internal

import (
	"fmt"

	"github.com/joshua468/team-budget-management/config"
	"github.com/joshua468/team-budget-management/models"
)

func AddExpense(expense models.Expense) error {
	if result := config.DB.Create(&expense); result.Error != nil {
		return fmt.Errorf("failed to create expense: %v", result.Error)
	}
	return nil
}

func GetExpensesByTeam(teamID uint) ([]models.Expense, error) {
	var expenses []models.Expense
	if result := config.DB.Where("team_id = ?", teamID).Find(&expenses); result.Error != nil {
		return nil, fmt.Errorf("failed to retrieve expenses: %v", result.Error)
	}
	return expenses, nil
}
