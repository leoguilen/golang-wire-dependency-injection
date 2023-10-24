package application

import (
	"sync"

	"github.com/google/wire"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/application/handlers"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/application/http"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/services"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/infra/data"
)

var (
	th     *handlers.HttpTaskHandler
	thOnce sync.Once

	ProviderSet wire.ProviderSet = wire.NewSet(
		ProvideHttpServer,
		ProvideHttpServerOptions,
		ProvideTaskHandler,
		domain.ProviderSet,
		data.ProviderSet,

		wire.Bind(new(handlers.ITaskHandler), new(*handlers.HttpTaskHandler)),
	)
)

func ProvideHttpServer(options http.HttpServerOptions, th handlers.ITaskHandler) *http.HttpServer {
	return &http.HttpServer{
		Options:     options,
		TaskHandler: th,
	}
}

func ProvideHttpServerOptions() http.HttpServerOptions {
	return http.HttpServerOptions{
		Port:   "8080",
		Prefix: "/api",
	}
}

func ProvideTaskHandler(ts services.ITaskService) *handlers.HttpTaskHandler {
	thOnce.Do(func() {
		th = &handlers.HttpTaskHandler{TaskService: ts}
	})
	return th
}
