package contracts

import "github.com/leoguilen/golang-wire-dependency-injection/internal/domain/models"

type TaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

func (req *TaskRequest) ToModel() *models.TaskWriteModel {
	return &models.TaskWriteModel{
		Title:       req.Title,
		Description: req.Description,
	}
}
