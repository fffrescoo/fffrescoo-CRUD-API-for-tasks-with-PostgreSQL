package tasksService

import (
	"gorm.io/gorm"
	"pedprojectFinal/internal/models"
)

// TaskRepository определяет интерфейс для работы с хранилищем задач
type TaskRepository interface {
	Create(task *models.Task) error
	GetByID(id uint) (*models.Task, error)
	Update(task *models.Task) error
	Delete(id uint) error
	Exists(id uint) (bool, error)
}

type taskRepository struct {
	db *gorm.DB
}

// NewTaskRepository создает новый экземпляр репозитория
func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) Create(task *models.Task) error {
	return r.db.Create(task).Error
}

func (r *taskRepository) GetByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.db.First(&task, id).Error
	return &task, err
}

func (r *taskRepository) Update(task *models.Task) error {
	return r.db.Save(task).Error
}

func (r *taskRepository) Delete(id uint) error {
	return r.db.Delete(&models.Task{}, id).Error
}

func (r *taskRepository) Exists(id uint) (bool, error) {
	var count int64
	err := r.db.Model(&models.Task{}).Where("id = ?", id).Count(&count).Error
	return count > 0, err
}
