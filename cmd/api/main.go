package main

import "github.com/leoguilen/golang-wire-dependency-injection/internal/application"

func main() {
	httpServer := application.WireHttpServer()
	panic(httpServer.Start())
}
