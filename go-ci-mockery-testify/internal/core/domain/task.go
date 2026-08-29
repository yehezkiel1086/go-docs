package domain

import "gorm.io/gorm"

const (
	StatusNotCompleted string = "not_completed"
	StatusInProgress string = "in_progress"
	StatusCompleted string = "completed"
)

type Task struct {
	gorm.Model

	Name string `json:"name" gorm:"not null;size:255"`
	Description string `json:"description" gorm:"not null;size:255"`
	Status string `json:"status" gorm:"not null;size:255;default:not_completed"`
}
