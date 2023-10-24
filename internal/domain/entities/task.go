package entities

import "github.com/google/uuid"

type Task struct {
	ID          uuid.UUID
	Title       string
	Description string
	IsDone      bool
}

func NewTask(title, description string) *Task {
	return &Task{
		ID:          uuid.New(),
		Title:       title,
		Description: description,
		IsDone:      false,
	}
}
