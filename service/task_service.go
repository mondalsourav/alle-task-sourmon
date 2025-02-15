package service

import (
	"alle-task/models"
	"alle-task/repository"
)

type TaskService struct {
	repo repository.TaskRepositoryInf
}

func NewTaskService(repo repository.TaskRepositoryInf) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task *models.Task) error {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks(status string, limit, offset int) ([]models.Task, error) {
	return s.repo.GetTasks(status, limit, offset)
}

func (s *TaskService) GetTaskByID(id uint) (*models.Task, error) {
	return s.repo.GetTaskByID(id)
}

func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.repo.UpdateTask(task)
}

func (s *TaskService) DeleteTask(id uint) error {
	return s.repo.DeleteTask(id)
}
