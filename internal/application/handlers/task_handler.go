package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/application/contracts"
	"github.com/leoguilen/golang-wire-dependency-injection/internal/domain/services"
)

type ITaskHandler interface {
	FindAll(w http.ResponseWriter, r *http.Request)
	FindById(w http.ResponseWriter, r *http.Request)
	Create(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type HttpTaskHandler struct {
	TaskService services.ITaskService
}

func NewTaskHandler(ts services.ITaskService) ITaskHandler {
	return &HttpTaskHandler{TaskService: ts}
}

func (th *HttpTaskHandler) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	var taskRequest contracts.TaskRequest

	err := json.NewDecoder(r.Body).Decode(&taskRequest)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(contracts.NewErrorResponse(err))
		return
	}

	task, err := th.TaskService.Create(r.Context(), taskRequest.ToModel())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(contracts.NewErrorResponse(err))
		return
	}

	w.Header().Set("Location", r.URL.Path+"/"+task.ID.String())
	w.WriteHeader(http.StatusCreated)
}

func (th *HttpTaskHandler) Delete(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	taskId, err := uuid.Parse(r.URL.Path[len("/api/tasks:delete/"):])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = th.TaskService.Delete(r.Context(), taskId)
	if err != nil {
		switch err.Error() {
		case "task not found":
			w.WriteHeader(http.StatusNotFound)
		default:
			{
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(contracts.NewErrorResponse(err))
			}
		}
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func (th *HttpTaskHandler) FindAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	tasks, err := th.TaskService.FindAll(r.Context())
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(contracts.NewErrorResponse(err))
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contracts.NewTaskResponseList(tasks))
}

func (th *HttpTaskHandler) FindById(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	taskId, err := uuid.Parse(r.URL.Path[len("/api/tasks/"):])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	task, err := th.TaskService.FindById(r.Context(), taskId)
	if err != nil {
		switch err.Error() {
		case "task not found":
			w.WriteHeader(http.StatusNotFound)
		default:
			{
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(contracts.NewErrorResponse(err))
			}
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contracts.NewTaskResponse(task))
}

func (th *HttpTaskHandler) Update(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPut {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	taskId, err := uuid.Parse(r.URL.Path[len("/api/tasks:update/"):])
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	var taskRequest contracts.TaskRequest

	err = json.NewDecoder(r.Body).Decode(&taskRequest)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(contracts.NewErrorResponse(err))
		return
	}

	task, err := th.TaskService.Update(r.Context(), taskId, taskRequest.ToModel())
	if err != nil {
		switch err.Error() {
		case "task not found":
			w.WriteHeader(http.StatusNotFound)
		default:
			{
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusInternalServerError)
				json.NewEncoder(w).Encode(contracts.NewErrorResponse(err))
			}
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(contracts.NewTaskResponse(task))
}
