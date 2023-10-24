package contracts

import (
	"github.com/google/uuid"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/models"
)

type TaskResponse struct {
	ID          uuid.UUID `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	IsDone      bool      `json:"is_done"`
}

type TaskResponseList struct {
	Tasks []TaskResponse `json:"data"`
}

func NewTaskResponse(tms *models.TaskReadModel) *TaskResponse {
	return &TaskResponse{
		ID:          tms.ID,
		Title:       tms.Title,
		Description: tms.Description,
		IsDone:      tms.IsDone,
	}
}

func NewTaskResponseList(tms []*models.TaskReadModel) *TaskResponseList {
	var tasks []TaskResponse = make([]TaskResponse, 0, len(tms))
	for _, tm := range tms {
		tasks = append(tasks, *NewTaskResponse(tm))
	}
	return &TaskResponseList{Tasks: tasks}
}
