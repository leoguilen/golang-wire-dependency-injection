// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package application

import (
	"github.com/leoguilen/golang-wire-dependency-injection/internal/application/http"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/infra/data"
)

// Injectors from wire.go:

func WireHttpServer() *http.HttpServer {
	httpServerOptions := ProvideHttpServerOptions()
	inMemoryTaskRepository := data.ProvideTaskRepository()
	taskService := domain.ProvideTaskService(inMemoryTaskRepository)
	httpTaskHandler := ProvideTaskHandler(taskService)
	httpServer := ProvideHttpServer(httpServerOptions, httpTaskHandler)
	return httpServer
}
