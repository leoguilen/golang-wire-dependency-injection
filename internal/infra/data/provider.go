package data

import (
	"sync"

	"github.com/google/wire"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/interfaces"
	data "github.com/leoguilen/golang-wire-dependency-injection/internal/infra/data/repositories"
)

var (
	tr     *data.InMemoryTaskRepository
	trOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideTaskRepository,
		wire.Bind(new(interfaces.ITaskRepository), new(*data.InMemoryTaskRepository)),
	)
)

func ProvideTaskRepository() *data.InMemoryTaskRepository {
	trOnce.Do(func() {
		tr = &data.InMemoryTaskRepository{}
	})
	return tr
}
