//go:build wireinject
// +build wireinject

package application

import (
	"github.com/google/wire"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/application/http"
)

func WireHttpServer() *http.HttpServer {
	panic(wire.Build(ProviderSet))
}
