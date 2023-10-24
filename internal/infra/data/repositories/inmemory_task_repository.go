package data

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/entities"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/interfaces"
)

var (
	_tasks map[uuid.UUID]*entities.Task = make(map[uuid.UUID]*entities.Task, 0)

	ErrTaskNotFound = errors.New("task not found")
)

type InMemoryTaskRepository struct{}

func NewInMemoryTaskRepository() interfaces.ITaskRepository {
	return InMemoryTaskRepository{}
}

func (InMemoryTaskRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if _, ok := _tasks[id]; !ok {
		return ErrTaskNotFound
	}
	delete(_tasks, id)
	return nil
}

func (InMemoryTaskRepository) FindAll(ctx context.Context) ([]*entities.Task, error) {
	tasks := make([]*entities.Task, 0, len(_tasks))
	for _, task := range _tasks {
		tasks = append(tasks, task)
	}
	return tasks, nil
}

func (InMemoryTaskRepository) FindById(ctx context.Context, id uuid.UUID) (*entities.Task, error) {
	task, ok := _tasks[id]
	if !ok {
		return nil, ErrTaskNotFound
	}
	return task, nil
}

func (InMemoryTaskRepository) Save(ctx context.Context, task *entities.Task) (*entities.Task, error) {
	if _, ok := _tasks[task.ID]; !ok {
		_tasks[task.ID] = task
		return task, nil
	}
	_tasks[task.ID] = task
	return task, nil
}
