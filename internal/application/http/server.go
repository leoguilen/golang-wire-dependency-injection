package http

import (
	"fmt"
	"log"
	"net/http"

	"github.com/leoguilen/golang-wire-dependency-injection/internal/application/handlers"
)

type HttpServer struct {
	Options     HttpServerOptions
	TaskHandler handlers.ITaskHandler
}

func NewHttpServer(options HttpServerOptions, th handlers.ITaskHandler) *HttpServer {
	return &HttpServer{
		Options:     options,
		TaskHandler: th,
	}
}

func (hs *HttpServer) Start() error {
	mux := http.NewServeMux()
	tasksPath := fmt.Sprintf("%s/tasks", hs.Options.Prefix)

	mux.HandleFunc(tasksPath, hs.TaskHandler.FindAll)
	mux.HandleFunc(fmt.Sprintf("%s/", tasksPath), hs.TaskHandler.FindById)
	mux.HandleFunc(fmt.Sprintf("%s:create", tasksPath), hs.TaskHandler.Create)
	mux.HandleFunc(fmt.Sprintf("%s:update/", tasksPath), hs.TaskHandler.Update)
	mux.HandleFunc(fmt.Sprintf("%s:delete/", tasksPath), hs.TaskHandler.Delete)

	log.Printf("Starting server on port %s", hs.Options.Port)
	return http.ListenAndServe(fmt.Sprintf(":%s", hs.Options.Port), mux)
}
