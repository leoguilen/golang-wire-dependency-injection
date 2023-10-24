package models

import (
	"github.com/google/uuid"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/entities"
)

type TaskReadModel struct {
	ID          uuid.UUID
	Title       string
	Description string
	IsDone      bool
}

type TaskWriteModel struct {
	Title       string
	Description string
}

func NewTaskReadModel(t *entities.Task) *TaskReadModel {
	return &TaskReadModel{
		ID:          t.ID,
		Title:       t.Title,
		Description: t.Description,
		IsDone:      t.IsDone,
	}
}
