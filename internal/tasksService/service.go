package tasksService

import (
	"errors"
	"pedprojectFinal/internal/models"
)

// TaskService определяет бизнес-логику
type TaskService interface {
	CreateTask(task string) (*models.Task, error)
	GetAll() ([]*models.Task, error)
	GetByID(id uint) (*models.Task, error)
	UpdateTask(id uint, newTask string) (*models.Task, error)
	DeleteTask(id uint) error
}

type taskService struct {
	repo TaskRepository
}

func NewTaskService(repo TaskRepository) TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(task string) (*models.Task, error) {
	if task == "" {
		return nil, errors.New("task cannot be empty")
	}

	newTask := &models.Task{
		Task: task,
	}

	if err := s.repo.Create(newTask); err != nil {
		return nil, err
	}

	return newTask, nil
}

func (s *taskService) GetAll() ([]*models.Task, error) {
	return s.repo.GetAll()
}

func (s *taskService) GetByID(id uint) (*models.Task, error) {
	return s.repo.GetByID(id)
}

func (s *taskService) UpdateTask(id uint, newTask string) (*models.Task, error) {
	if newTask == "" {
		return nil, errors.New("task cannot be empty")
	}

	task, err := s.repo.GetByID(id)
	if err != nil {
		return nil, err
	}

	task.Task = newTask
	if err := s.repo.Update(task); err != nil {
		return nil, err
	}

	return task, nil
}

func (s *taskService) DeleteTask(id uint) error {
	return s.repo.Delete(id)
}
