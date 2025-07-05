package handlers

import (
	"context"
	"fmt"
	"pedprojectFinal/internal/tasksService"
	"pedprojectFinal/internal/web/tasks"
)

type Handler struct {
	service tasksService.TaskService
}

func NewHandler(service tasksService.TaskService) *Handler {
	return &Handler{service: service}
}

// GET /tasks
func (h *Handler) GetTasks(ctx context.Context, _ tasks.GetTasksRequestObject) (tasks.GetTasksResponseObject, error) {
	allTasks, err := h.service.GetAll()
	if err != nil {
		return nil, err
	}

	// Инициализируем пустой слайс, чтобы при пустом списке возвращался []
	response := make(tasks.GetTasks200JSONResponse, 0)
	for _, t := range allTasks {
		task := tasks.Task{
			Id:   uintPtr(uint(t.ID)),
			Task: &t.Task,
		}
		response = append(response, task)
	}

	return response, nil
}

// POST /tasks
func (h *Handler) PostTasks(ctx context.Context, req tasks.PostTasksRequestObject) (tasks.PostTasksResponseObject, error) {
	if req.Body == nil || req.Body.Task == nil {
		return nil, fmt.Errorf("invalid request body: missing task")
	}

	created, err := h.service.CreateTask(*req.Body.Task)
	if err != nil {
		return nil, err
	}

	resp := tasks.PostTasks201JSONResponse{
		Id:   uintPtr(uint(created.ID)),
		Task: &created.Task,
	}

	return resp, nil
}

// PATCH /tasks/{id}
func (h *Handler) PatchTasksId(ctx context.Context, req tasks.PatchTasksIdRequestObject) (tasks.PatchTasksIdResponseObject, error) {
	id := uint(req.Id)

	if req.Body == nil || req.Body.Task == nil {
		return nil, fmt.Errorf("missing updated task text")
	}

	updated, err := h.service.UpdateTask(id, *req.Body.Task)
	if err != nil {
		return tasks.PatchTasksId404Response{}, nil // Task not found
	}

	resp := tasks.Task{
		Id:   uintPtr(uint(updated.ID)),
		Task: &updated.Task,
	}

	return tasks.PatchTasksId200JSONResponse(resp), nil
}

// DELETE /tasks/{id}
func (h *Handler) DeleteTasksId(ctx context.Context, req tasks.DeleteTasksIdRequestObject) (tasks.DeleteTasksIdResponseObject, error) {
	id := uint(req.Id)

	err := h.service.DeleteTask(id)
	if err != nil {
		return tasks.DeleteTasksId404Response{}, nil // Task not found
	}

	return tasks.DeleteTasksId204Response{}, nil
}

// вспомогательная функция для uint
func uintPtr(v uint) *uint {
	return &v
}
