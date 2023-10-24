package interfaces

import (
	"context"

	"github.com/google/uuid"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/entities"
)

type ITaskRepository interface {
	FindAll(ctx context.Context) ([]*entities.Task, error)
	FindById(ctx context.Context, id uuid.UUID) (*entities.Task, error)
	Save(ctx context.Context, task *entities.Task) (*entities.Task, error)
	Delete(ctx context.Context, id uuid.UUID) error
}
