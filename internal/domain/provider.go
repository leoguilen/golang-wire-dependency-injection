package domain

import (
	"sync"

	"github.com/google/wire"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/interfaces"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/services"
)

var (
	ts     *services.TaskService
	tsOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideTaskService,
		wire.Bind(new(services.ITaskService), new(*services.TaskService)),
	)
)

func ProvideTaskService(tr interfaces.ITaskRepository) *services.TaskService {
	tsOnce.Do(func() {
		ts = &services.TaskService{TaskRepository: tr}
	})
	return ts
}
