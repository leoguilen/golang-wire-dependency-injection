package services

import (
	"context"

	"github.com/google/uuid"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/entities"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/interfaces"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/models"
)

type ITaskService interface {
	FindAll(ctx context.Context) ([]*models.TaskReadModel, error)
	FindById(ctx context.Context, id uuid.UUID) (*models.TaskReadModel, error)
	Create(ctx context.Context, t *models.TaskWriteModel) (*models.TaskReadModel, error)
	Update(ctx context.Context, id uuid.UUID, t *models.TaskWriteModel) (*models.TaskReadModel, error)
	Delete(ctx context.Context, id uuid.UUID) error
}

type TaskService struct {
	TaskRepository interfaces.ITaskRepository
}

func NewTaskService(taskRepository interfaces.ITaskRepository) ITaskService {
	return TaskService{
		TaskRepository: taskRepository,
	}
}

func (ts TaskService) Delete(ctx context.Context, id uuid.UUID) error {
	return ts.TaskRepository.Delete(ctx, id)
}

func (ts TaskService) FindAll(ctx context.Context) ([]*models.TaskReadModel, error) {
	tasks, err := ts.TaskRepository.FindAll(ctx)
	if err != nil {
		return nil, err
	}

	tasksModel := make([]*models.TaskReadModel, 0, len(tasks))
	for _, task := range tasks {
		tasksModel = append(tasksModel, models.NewTaskReadModel(task))
	}

	return tasksModel, nil
}

func (ts TaskService) FindById(ctx context.Context, id uuid.UUID) (*models.TaskReadModel, error) {
	task, err := ts.TaskRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.NewTaskReadModel(task), nil
}

func (ts TaskService) Update(ctx context.Context, id uuid.UUID, t *models.TaskWriteModel) (*models.TaskReadModel, error) {
	task, err := ts.TaskRepository.FindById(ctx, id)
	if err != nil {
		return nil, err
	}

	task.Title = t.Title
	task.Description = t.Description

	task, err = ts.TaskRepository.Save(ctx, task)
	if err != nil {
		return nil, err
	}

	return models.NewTaskReadModel(task), nil
}

func (ts TaskService) Create(ctx context.Context, t *models.TaskWriteModel) (*models.TaskReadModel, error) {
	task := entities.NewTask(t.Title, t.Description)

	insertedTask, err := ts.TaskRepository.Save(ctx, task)
	if err != nil {
		return nil, err
	}

	return models.NewTaskReadModel(insertedTask), nil
}
