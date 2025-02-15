package repository

import (
	"alle-task/models"
	"gorm.io/gorm"
)

type TaskRepositoryInf interface {
	CreateTask(task *models.Task) error
	GetTasks(status string, limit, offset int) ([]models.Task, error)
	GetTaskByID(id uint) (*models.Task, error)
	UpdateTask(task *models.Task) error
	DeleteTask(id uint) error
}

type TaskRepository struct {
	DB *gorm.DB
}

func NewTaskRepository(db *gorm.DB) *TaskRepository {
	return &TaskRepository{DB: db}
}

func (r *TaskRepository) CreateTask(task *models.Task) error {
	return r.DB.Create(task).Error
}

func (r *TaskRepository) GetTasks(status string, limit, offset int) ([]models.Task, error) {
	var tasks []models.Task
	query := r.DB.Model(&models.Task{})

	if status != "" {
		query = query.Where("status = ?", status)
	}

	err := query.Limit(limit).Offset(offset).Find(&tasks).Error
	return tasks, err
}

func (r *TaskRepository) GetTaskByID(id uint) (*models.Task, error) {
	var task models.Task
	err := r.DB.First(&task, id).Error
	return &task, err
}

func (r *TaskRepository) UpdateTask(task *models.Task) error {
	return r.DB.Save(task).Error
}

func (r *TaskRepository) DeleteTask(id uint) error {
	return r.DB.Delete(&models.Task{}, id).Error
}
