package models

import (
	"gorm.io/gorm"
)

type Team struct {
	gorm.Model
	Name     string    `json:"name"`
	Budget   float32   `json:"budget"`
	Expenses []Expense `gorm:"foreignKey:TeamID" json:"expenses"`
}

type Expense struct {
	gorm.Model
	TeamID   uint    `json:"team_id"`
	Amount   float64 `json:"amount"`
	Approved bool    `json:"approved"`
	Reason   string  `json:"reason"`
}
